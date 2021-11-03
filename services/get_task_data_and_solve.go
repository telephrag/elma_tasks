package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mservice/config"
	"mservice/models"
	"mservice/solver_wrappers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetTaskDataAndSolve() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	rootCtx := context.Background()

	r.Get("/tasks/{name}", func(rw http.ResponseWriter, r *http.Request) {
		taskName := chi.URLParam(r, "name")
		// TODO: Remove this and handle "task/" in diffrent service
		if taskName == "" {
			stdInternalServerError(&rw, nil, "No task name provided.")
			return
		}

		resp, err := http.Get("http://" + config.DataProviderAddr + "/tasks/" + taskName)
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn get data from provider.")
			return
		}

		if resp.StatusCode != http.StatusOK {
			rw.WriteHeader(resp.StatusCode)
			return
		}

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn't read response contents.")
			return
		}
		rw.Write(append([]byte("Data received succesfully\n"), content...))
		resp.Body.Close()

		var data [][]interface{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			stdInternalServerError(&rw, err, "")
			return
		}

		// Consider adding timeout to context
		taskCtx := context.WithValue(rootCtx, "name", taskName)
		rj, err := solver_wrappers.SelectWrapper(data, taskCtx)
		if err != nil {
			stdInternalServerError(&rw, err, "Error occured during solving.")
			return
		}

		rw.Header().Set("Content-type", "application/json")

		req, err := json.Marshal(rj)
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn't package request data.")
			return
		}

		resp, err = http.Post("http://"+config.DataProviderAddr+"/tasks/solution", "application/json", bytes.NewBuffer(req))
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn't send data to provider")
		}

		if resp.StatusCode != http.StatusOK {
			rw.WriteHeader(resp.StatusCode)
			return
		}

		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn't read response contents.")
			return
		}
		rw.Write(append([]byte("Response received succesfully\n"), content...))

		var res models.Response
		err = json.Unmarshal(content, &res)
		if err != nil {
			stdInternalServerError(&rw, err, "Couldn't unmarshal response from the server.")
		}
	})

	// TODO: Try using middleware as shown here: https://go-chi.io/#/pages/middleware
	// TODO: Try middleware chains as shown here: https://stackoverflow.com/questions/49025811/http-handler-function
	return r
}

func stdInternalServerError(rw *http.ResponseWriter, err error, msg string) {
	fmt.Println(msg)
	// (*rw).Write([]byte(err.Error() + "\n" + msg))
	// (*rw).WriteHeader(http.StatusInternalServerError)
}

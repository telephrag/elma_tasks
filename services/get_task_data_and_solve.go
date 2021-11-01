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

func GetTaskDataAndSolveService() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	var rj models.ResultJson
	var isErr bool
	var taskName string
	rootCtx := context.Background()

	r.Group(func(r chi.Router) {
		r.Get("/tasks/{name}", func(rw http.ResponseWriter, r *http.Request) {
			var err error
			defer func() {
				if err != nil {
					isErr = true
				}
			}()

			taskName = chi.URLParam(r, "name")
			// TODO: Remove this and handle "task/" in diffrent service
			if taskName == "" {
				stdInternalServerError(&rw, err, "No task name provided.")
				return
			}

			resp, err := http.Get("http://" + config.DataProviderAddr + "/tasks/" + taskName)
			if err != nil {
				stdInternalServerError(&rw, err, "Couldn get data from provider.")
				return
			}

			if resp.StatusCode != http.StatusOK {
				stdInternalServerError(&rw, err, "Bad response from data provider.")
				return
			}

			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				stdInternalServerError(&rw, err, "Couldn't read response contents.")
				return
			}
			rw.Write(append([]byte("Data received succesfully\n"), content...))

			var data [][]interface{}
			err = json.Unmarshal(content, &data)
			if err != nil {
				stdInternalServerError(&rw, err, "")
				return
			}

			taskCtx := context.WithValue(rootCtx, "name", taskName)
			rr, err := solver_wrappers.SelectWrapper(data, taskCtx)
			if err != nil {
				stdInternalServerError(&rw, err, "Error occured during solving.")
				return
			}

			rj, err = models.NewResultJson(rr, data, taskCtx)
			if err != nil {
				stdInternalServerError(&rw, err, "Error occured during response packaging.")
				return
			}
		})

		if isErr {
			fmt.Println("Error occured in handlerFn.")
			return
		}

		r.Post("/tasks/"+taskName, func(rw http.ResponseWriter, r *http.Request) {
			req, err := json.Marshal(rj)
			if err != nil {
				stdInternalServerError(&rw, err, "Couldn't package request data.")
				return
			}

			resp, err := http.Post("http://"+config.DataProviderAddr+"/tasks/solution", "application/json", bytes.NewBuffer(req))
			if err != nil {
				stdInternalServerError(&rw, err, "Couldn't send data to provider")
			}

			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				stdInternalServerError(&rw, err, "Couldn't read response contents.")
				return
			}
			rw.Write(append([]byte("Response received succesfully\n"), content...))
		})
	})

	return r
}

func stdInternalServerError(rw *http.ResponseWriter, err error, msg string) {
	(*rw).Write([]byte(err.Error() + "\n" + msg))
	(*rw).WriteHeader(http.StatusInternalServerError)
}

package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mservice/config"
	"mservice/solver_wrappers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetTaskDataAndSolveService() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get(
		"/tasks/{name}",
		func(rw http.ResponseWriter, r *http.Request) {
			taskName := chi.URLParam(r, "name")

			if taskName == "" {
				fmt.Println("No task name received.")
				rw.WriteHeader(http.StatusNotFound)
				return
			}

			resp, err := http.Get("http://" + config.DataProviderAddr + "/tasks/" + taskName)
			if err != nil {
				rw.Write([]byte(err.Error() + "\nCouldn get data from provider."))
				rw.WriteHeader(resp.StatusCode)
				return
			}

			if resp.StatusCode != http.StatusOK {
				rw.Write([]byte(err.Error() + "\nBad response from provider."))
				rw.WriteHeader(resp.StatusCode)
				return
			}

			// should I atomically write to solution body
			// while solving multiple tasks asynchronously?
			// maybe later

			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				rw.Write([]byte(err.Error()))
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.Write(append([]byte("Data received succesfully\n"), content...))

			var data [][]interface{}
			err = json.Unmarshal(content, &data)
			if err != nil {
				rw.Write([]byte(err.Error()))
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			var resA [][][]float64
			var resK []float64
			rootCtx := context.Background()
			taskCtx := context.WithValue(rootCtx, "name", taskName)
			go solver_wrappers.SelectWrapper(data, taskCtx) // TODO: Reuse content when sending response to server
		},
	)

	return r
}

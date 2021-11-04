package public

import (
	"bytes"
	"context"
	"io/ioutil"
	"mservice/config"
	"mservice/util"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetDataAndSolve() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tasks/{name}", func(rw http.ResponseWriter, r *http.Request) {

		var err error = nil
		var errMsg string = ""
		defer util.Ordinary500(rw, err, errMsg)

		taskName := chi.URLParam(r, "name")
		if taskName == "" {
			return
		}

		resp, err := http.Get(
			config.ProtoHttp + config.LocalInternalAddr + "/tasks/" + taskName,
		)

		if err != nil {
			return
		}

		if resp.StatusCode != http.StatusOK {
			rw.WriteHeader(resp.StatusCode)
		}

		content, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return
		}

		taskCtx := context.WithValue(r.Context(), "taskName", taskName)
		postReq, err := http.NewRequestWithContext(
			taskCtx,
			"POST",
			config.ProtoHttp+config.LocalInternalAddr+"/tasks/solution",
			bytes.NewBuffer(content),
		)
		postReq.Header.Set("Content-Type", "application/json")
		if err != nil {
			errMsg = "Failed to form a request.\n"
			return
		}

		postReq.Header.Set("taskName", taskName)

		client := http.DefaultClient
		resp, err = client.Do(postReq)

		if err != nil {
			return
		}

		if resp.StatusCode != http.StatusOK {
			rw.WriteHeader(resp.StatusCode)
			return
		}

		content, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return
		}
		rw.Write(content)
	})

	return r
}

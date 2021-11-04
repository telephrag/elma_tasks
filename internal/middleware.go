package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"mservice/solver_wrappers"
	"mservice/util"
	"net/http"
)

func Solver(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		var errMsg string = ""
		var err error = nil
		defer util.Ordinary500(rw, err, errMsg)

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg = "Failed to read from response body.\n"
			return
		}
		r.Body.Close()

		var data [][]interface{}
		err = json.Unmarshal(reqBody, &data)
		if err != nil {
			errMsg = "Failed to unmarshal data.\n"
			return
		}

		taskName := r.Header.Get("taskName")
		rj, err := solver_wrappers.SelectWrapper(data, taskName)
		if err != nil {
			errMsg = "Failed to solve using given data.\n"
			return
		}

		ctxWithResult := context.WithValue(r.Context(), "result", rj)

		next.ServeHTTP(rw, r.Clone(ctxWithResult))
	})
}

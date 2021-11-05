package localmw

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"mservice/models"
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

		var tasks []models.Task
		err = json.Unmarshal(reqBody, &tasks)
		if err != nil {
			errMsg = "Failed to unmarshal data.\n"
			return
		}

		res, err := solver_wrappers.SelectWrapper(tasks)
		if err != nil {
			errMsg = "Failed to solve using given data.\n"
			return
		}

		ctxWithResult := context.WithValue(r.Context(), "results", res)

		next.ServeHTTP(rw, r.Clone(ctxWithResult))
	})
}

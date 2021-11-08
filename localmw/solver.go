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

		var err error
		code := 0
		defer util.HttpErrWriter(rw, err, code)

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		r.Body.Close()

		var tasks []models.Task
		err = json.Unmarshal(reqBody, &tasks)
		if err != nil {
			return
		}

		res, err := solver_wrappers.SelectWrapper(tasks)
		if err != nil {
			return
		}

		ctxWithResult := context.WithValue(r.Context(), "results", res)

		next.ServeHTTP(rw, r.Clone(ctxWithResult))
	})
}

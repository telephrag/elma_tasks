package localmw

import (
	"errors"
	"fmt"
	"mservice/config"
	"mservice/models"
	"mservice/util"
	"net/http"
)

func Printer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var err error = nil
		code := 0
		defer util.HttpErrWriter(rw, err, code)

		res, ok := r.Context().Value("results").([]models.Result)
		if !ok {
			err = errors.New(config.CtxWrongType)
			return
		}

		for i := range res {
			PrintResult(rw, res[i])
		}

		next.ServeHTTP(rw, r)
	})
}

func PrintResult(rw http.ResponseWriter, res models.Result) {
	rw.Write([]byte(res.TaskName + "\n"))
	rw.Write([]byte(fmt.Sprintln(res.Results.Results...)))
}

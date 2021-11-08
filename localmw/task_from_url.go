package localmw

import (
	"context"
	"errors"
	"mservice/config"
	"mservice/util"
	"net/http"

	"github.com/go-chi/chi"
)

func TaskFromURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var err error
		var code int
		defer util.HttpErrWriter(rw, err, code)

		taskName := chi.URLParam(r, "name")
		if taskName == "" {
			err = errors.New(config.NoTask)
			return
		}

		tasks := make([]string, 1)
		tasks[0] = taskName

		ctx := context.WithValue(r.Context(), "tasks", tasks)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

package localmw

import (
	"context"
	"encoding/json"
	"mservice/util"
	"net/http"
)

func TaskParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var err error
		var code int
		defer util.HttpErrWriter(rw, err, code)

		t := r.Header.Get("tasks")
		var tasks []string
		err = json.Unmarshal([]byte(t), &tasks)
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), "tasks", tasks)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

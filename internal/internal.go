package internal

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InternalService() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/{name}", GetTaskData)
		r.With(Solver).Post("/solution", PostBack)
	})

	return r
}

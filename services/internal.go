package services

import (
	"mservice/handlerfuncs"
	"mservice/localmw"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InternalService() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/{name}", handlerfuncs.GetTaskData)
		r.With(localmw.Solver, localmw.Printer).Post("/solution", handlerfuncs.PostBack)
	})

	return r
}

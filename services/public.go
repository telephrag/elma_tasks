package services

import (
	"mservice/handlerfuncs"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetDataAndSolveService() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/{name}", handlerfuncs.GetDataAndSolve)
		r.Get("/", handlerfuncs.GetDataAndSolve)
	})

	return r
}

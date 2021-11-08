package services

import (
	"mservice/handlerfuncs"
	"mservice/localmw"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetDataAndSolveService() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.With(localmw.TaskFromURL).Get("/{name}", handlerfuncs.GetDataAndSolve)
		r.With(localmw.TaskParser).Get("/", handlerfuncs.GetDataAndSolve)
	})

	return r
}

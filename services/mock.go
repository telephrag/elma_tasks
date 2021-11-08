package services

import (
	"mservice/handlerfuncs"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func MockService() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tasks/", handlerfuncs.Mock)

	return r
}

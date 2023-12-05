package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/user", handleUserPost)
	r.Get("/user", handleUserGet)
	return r
}

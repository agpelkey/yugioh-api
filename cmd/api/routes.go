package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// create mux server
	mux := chi.NewRouter()

	mux.Get("/", app.Home)

	return mux
}

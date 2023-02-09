package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// create mux server
	mux := chi.NewRouter()

	mux.Get("/", app.Home)
	mux.Get("/cards", app.GetAllCards)
	mux.Get("/cards/{id}", app.GetCardByID)
	mux.Get("/cards/level/{level}", app.GetCardByLevel)
	// mux.Get("/cards/attack/{attack}", app.GetCardsByAttack) // I need to figure out the routing to search for a lower and upper attack limit
	mux.Post("/cards/add", app.InsertNewCard)
	// mux.Put("/cards/update") for updating card - although I dont think a user would ever update a card unless Konami changed something
	mux.Delete("/cards/delete/{id}", app.DeleteCardWithID)
	return mux
}

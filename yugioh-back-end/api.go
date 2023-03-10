package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	listenAddr string
	db         DatabaseRepo
}

// function to create a new API server
func NewAPIServer(listenAddr string, db DatabaseRepo) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

// function to be called in main to run the server.
// includes the routes for the api as well
func (s *APIServer) Run() {
	mux := chi.NewRouter()

	mux.Get("/", s.Home)
	mux.Get("/cards", s.GetAllCards)
	mux.Get("/cards/{id}", s.GetCardByID)
	mux.Get("/cards/level/{level}", s.GetCardByLevel)
	// mux.Get("/cards/attack/{attack}", app.GetCardsByAttack) // I need to figure out the routing to search for a lower and upper attack limit
	mux.Post("/cards/add", s.InsertNewCard)
	// mux.Put("/cards/update") for updating card - although I dont think a user would ever update a card unless Konami changed something
	mux.Delete("/cards/delete/{id}", s.DeleteCardWithID)

	mux.("/account", s.HandleAccount)

	log.Println("starting server on port ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)
}

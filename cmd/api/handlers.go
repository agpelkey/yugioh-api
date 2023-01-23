package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// Function to get all cards from the DB
func (app *application) GetAllCards(w http.ResponseWriter, r *http.Request) {
	c, err := app.DB.AllCards()
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, c)
}

// Function to get a single card by ID
func (app *application) GetCardByID(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "id")
	nameID, err := strconv.Atoi(name)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload, err := app.DB.OneCardByID(nameID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

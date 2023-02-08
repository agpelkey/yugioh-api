package main

import (
	"net/http"
	"strconv"
	"yugioh-api/internal/models"

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

// Function to get card by level
func (app *application) GetCardByLevel(w http.ResponseWriter, r *http.Request) {
	level := chi.URLParam(r, "level")
	levelNum, err := strconv.Atoi(level)
	if err != nil {
		app.errorJSON(w, err)
	}

	payload, err := app.DB.GetCardByLevel(levelNum)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

// I wanted this function to return a range of attack i.e. 2000-3000
// havent been able to figure it out yet, will come back to it.
func (app *application) GetCardsByAttack(w http.ResponseWriter, r *http.Request) {
	attck := chi.URLParam(r, "attack")
	attckNum, err := strconv.Atoi(attck)
	if err != nil {
		app.errorJSON(w, err)
	}

	payload, err := app.DB.GetCardsByAttack(attckNum)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) InsertNewCard(w http.ResponseWriter, r *http.Request) {
	var card models.YugiohCard

	err := app.readJSON(w, r, &card)
	if err != nil {
		app.errorJSON(w, err)
	}

	newID, err := app.DB.AddNewCard(card)
	if err != nil {
		app.errorJSON(w, err)
	}

	app.writeJSON(w, http.StatusAccepted, newID)
}

func (app *application) DeleteCardWithID(w http.ResponseWriter, r *http.Request) {
	delete_card := chi.URLParam(r, "id")
	cardID, err := strconv.Atoi(delete_card)
	if err != nil {
		app.errorJSON(w, err)
	}

	err = app.DB.DeleteCard(cardID)
	if err != nil {
		app.errorJSON(w, err)
	}

	resp := JSONResponse{
		Error:   false,
		Message: "movie deleted",
	}

	app.writeJSON(w, http.StatusOK, resp)

}

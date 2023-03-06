package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *APIServer) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Yugioh app up and running",
		Version: "1.0.0",
	}

	_ = s.writeJSON(w, http.StatusOK, payload)
}

// Function to get all cards from the DB
func (s *APIServer) GetAllCards(w http.ResponseWriter, r *http.Request) {
	c, err := s.db.AllCards()
	if err != nil {
		s.errorJSON(w, err)
	}

	_ = s.writeJSON(w, http.StatusOK, c)
}

// Function to get a single card by ID
func (s *APIServer) GetCardByID(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "id")
	nameID, err := strconv.Atoi(name)
	if err != nil {
		s.errorJSON(w, err)
		return
	}

	payload, err := s.db.OneCardByID(nameID)
	if err != nil {
		s.errorJSON(w, err)
		return
	}

	_ = s.writeJSON(w, http.StatusOK, payload)
}

// Function to get card by level
func (s *APIServer) GetCardByLevel(w http.ResponseWriter, r *http.Request) {
	level := chi.URLParam(r, "level")
	levelNum, err := strconv.Atoi(level)
	if err != nil {
		s.errorJSON(w, err)
	}

	payload, err := s.db.GetCardByLevel(levelNum)
	if err != nil {
		s.errorJSON(w, err)
	}

	_ = s.writeJSON(w, http.StatusOK, payload)

}

// I wanted this function to return a range of attack i.e. 2000-3000
// havent been able to figure it out yet, will come back to it.
func (s *APIServer) GetCardsByAttack(w http.ResponseWriter, r *http.Request) {
	attck := chi.URLParam(r, "attack")
	attckNum, err := strconv.Atoi(attck)
	if err != nil {
		s.errorJSON(w, err)
	}

	payload, err := s.db.GetCardsByAttack(attckNum)
	if err != nil {
		s.errorJSON(w, err)
	}

	_ = s.writeJSON(w, http.StatusOK, payload)
}

func (s *APIServer) InsertNewCard(w http.ResponseWriter, r *http.Request) {
	var card YugiohCard

	err := s.readJSON(w, r, &card)
	if err != nil {
		s.errorJSON(w, err)
	}

	newID := s.db.AddNewCard(&card)
	if err != nil {
		s.errorJSON(w, err)
	}

	s.writeJSON(w, http.StatusAccepted, newID)
}

func (s *APIServer) DeleteCardWithID(w http.ResponseWriter, r *http.Request) {
	delete_card := chi.URLParam(r, "id")
	cardID, err := strconv.Atoi(delete_card)
	if err != nil {
		s.errorJSON(w, err)
	}

	err = s.db.DeleteCard(cardID)
	if err != nil {
		s.errorJSON(w, err)
	}

	resp := JSONResponse{
		Error:   false,
		Message: "card deleted",
	}

	s.writeJSON(w, http.StatusOK, resp)

}

/*
func (app *application) UpdateCardWithID(w http.ResponseWriter, r *http.Request) {
	var payload models.YugiohCard

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
	}

	card, err := app.DB.OneCardByID(payload.ID)
	if err != nil {
		app.errorJSON(w, err)
	}

	card.Name = payload.Name
	card.Level = payload.Level
	card. = payload.Attack


}
*/

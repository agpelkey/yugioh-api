package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Struct to hold card information
type YugiohCard struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Level   string `json:"level"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

// struct to hold user account information
type Account struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

// struct that holds a new account request information.
// will be passed to the handler responsible for creating new user accounts.
type CreateAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// function for new account request. Might not be the best place to put this function.
func NewAccount(username, email, password string) (*Account, error) {
	// encrypt the password
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		Username:          username,
		Email:             email,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC()}, nil
}

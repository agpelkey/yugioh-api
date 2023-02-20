package main

import "time"

// Struct to hold card information
type YugiohCard struct {
	Name    string `json:"name"`
	Level   string `json:"level"`
	Attack  string `json:"attack"`
	Defense string `json:"defense"`
	ID      int    `json:"id"`
}

type Account struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

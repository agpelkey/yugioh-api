package main

import "time"

// Struct to hold card information
type YugiohCard struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Level   string `json:"level"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

type Account struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

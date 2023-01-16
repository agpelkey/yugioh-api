package models

import "time"

// Struct to hold card information
type YugiohCard struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}


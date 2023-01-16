package models

// Struct to hold card information
type YugiohCard struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

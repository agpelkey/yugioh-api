package models

// Struct to hold card information
type YugiohCard struct {
	Name    string `json:"name"`
	Level   string `json:"level"`
	Attack  string `json:"attack"`
	Defense string  `json:"defense"`
	ID      int    `json:"id"`
}

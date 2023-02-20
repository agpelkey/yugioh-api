package main

import (
	"database/sql"
)

// Database interface
type DatabaseRepo interface {
	Connection() *sql.DB
	AllCards() ([]*YugiohCard, error)
	OneCardByID(id int) ([]*YugiohCard, error)
	GetCardByLevel(level int) ([]*YugiohCard, error)
	GetCardsByAttack(attack int) ([]*YugiohCard, error)
	AddNewCard(card YugiohCard) (int, error)
	DeleteCard(id int) error
	UpdateCard(card YugiohCard) error
}

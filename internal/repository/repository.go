package repository

import (
	"database/sql"
	"yugioh-api/internal/models"
)

// Database interface
type DatabaseRepo interface {
	Connection() *sql.DB
	AllCards() ([]*models.YugiohCard, error)
	OneCardByID(id int) ([]*models.YugiohCard, error)
	GetCardByLevel(level int) ([]*models.YugiohCard, error)
	GetCardsByAttack(attack int) ([]*models.YugiohCard, error)
	AddNewCard(card models.YugiohCard) (int, error)
	DeleteCard(id int) error 
}

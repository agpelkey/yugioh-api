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
}

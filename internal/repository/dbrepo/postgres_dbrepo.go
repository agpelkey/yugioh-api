package dbrepo

import (
	"context"
	"database/sql"
	"time"
	"yugioh-api/internal/models"
)

// create connection pool to database
type PostgresDBRepo struct {
	DB *sql.DB // DB = database connection pool
}

// time to interact with DB before timeout
const dbtimeout = time.Second * 3

// function to connect to DB
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

// Function to return all cards
func (m *PostgresDBRepo) AllCards() ([]*models.YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Connect to db and get a list of all cards
	query := `SELECT * FROM yugioh_cards`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create new variable that is pointer to yugioh card struct
	var cards []*models.YugiohCard

	for rows.Next() {
		var card models.YugiohCard
		err := rows.Scan(
			&card.Name,
			&card.Level,
			&card.Attack,
			&card.Defense,
		)
		if err != nil {
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, nil

}

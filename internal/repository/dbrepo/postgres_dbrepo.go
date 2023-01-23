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
			&card.ID,
		)
		if err != nil {
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, nil

}

// Function to get card by name
func (m *PostgresDBRepo) OneCardByName(id int) ([]*models.YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card by id
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE id = $1`

	// Executes the query that returns rows
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create new variable which is of type slice of pointer to models.yugiohcards to hold data 
	var cards []*models.YugiohCard

	// Create variable 'card' pointer to models.YugiohCard then iterate through rows data, Scanning into &card
	for rows.Next() {
		var card models.YugiohCard
		err := rows.Scan(
			&card.ID,
			&card.Name,
			&card.Level,
			&card.Attack,
			&card.Defense,
		)
		if err != nil {
			return nil, err
		}

		// add scanned data back into cards slice
		cards = append(cards, &card)
	}

	return cards, nil

}

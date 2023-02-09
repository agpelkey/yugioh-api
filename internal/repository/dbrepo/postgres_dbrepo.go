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

// Function to get card by id
func (m *PostgresDBRepo) OneCardByID(id int) ([]*models.YugiohCard, error) {
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

func (m *PostgresDBRepo) GetCardByLevel(level int) ([]*models.YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card by id
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE level = $1`

	// Exectues the query that returns the rows
	rows, err := m.DB.QueryContext(ctx, query, level)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create new variable which is of type slice of pointer to models.yugiohcard
	var cards []*models.YugiohCard

	// Create variable 'card' pointer to models.Yugiohcard then iterare through rows data,
	// scanning into &card

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

func (m *PostgresDBRepo) GetCardsByAttack(attack int) ([]*models.YugiohCard, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card with attack between desired range
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE attack BETWEEN $1 and $2`

	// Execute the query that returns the rows
	rows, err := m.DB.QueryContext(ctx, query, attack)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create variable which is of type slice pointer to models.yugiohcard
	var cards []*models.YugiohCard

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

		cards = append(cards, &card)
	}

	return cards, nil
}

func (m *PostgresDBRepo) AddNewCard(card models.YugiohCard) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `insert into yugioh_cards (name, level, attack, defense) values ($1, $2, $3, $4) returning id`

	var newID int

	err := m.DB.QueryRowContext(ctx, query,
		card.Name,
		card.Level,
		card.Attack,
		card.Defense,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil

}

func (m *PostgresDBRepo) DeleteCard(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	stmt := `delete from yugioh_cards where id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) UpdateCard(card models.YugiohCard) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `update yugioh_cards set name = $1, level = $2, attack = $3, defense = $4 where id = $5`

	_, err := m.DB.ExecContext(ctx, query,
		card.Name,
		card.Level,
		card.Attack,
		card.Defense,
	)
	if err != nil {
		return err
	}

	return nil

}

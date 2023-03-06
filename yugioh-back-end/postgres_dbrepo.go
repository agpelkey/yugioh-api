package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
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

// create connection pool to database
type PostgresDBRepo struct {
	db *sql.DB // DB = database connection pool
}

// time to interact with DB before timeout
const dbtimeout = time.Second * 3

// function to connect to DB
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.db
}

// function to create new Postgres DB from docker image
func NewPostgresStore() (*PostgresDBRepo, error) {
	connstr := "user=postgres dbname=postgres password=yugioh sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDBRepo{
		db: db,
	}, nil
}

// function to create DB tables
func (m *PostgresDBRepo) Init() error {
	return m.createAccountTable()
	return m.createMonsterCardTable()
}

// function to create cards table in DB
func (m *PostgresDBRepo) createMonsterCardTable() error {
	query := `CREATE TABLE IF NOT EXISTS yugioh_monster_cards (
		id serial primary key,
		name varchar(50) not null,
		level varchar(50) not null,
		attack int not null,
		defense int not null
	);`

	_, err := m.db.Exec(query)
	return err
}

// function to create user table in DB
func (m *PostgresDBRepo) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS yugioh_accounts 
				(
					id serial primary key,
					username varchar(50),
					first_name varchar(50),
					last_name varchar(50),
					email varchar(50),
					encrypted_password varchar(100),
					created_at timestamp
				);`

	_, err := m.db.Exec(query)
	return err
}

func (m *PostgresDBRepo) CreateAccount(acc *Account) error {
	query := `INSERT INTO  yugioh_accounts
			  (username, first_name, last_name, email, encrypted_password, created_at)
			  values
			  ($1, $2, $3, $4, $5, $6)`

	_, err := m.db.Query(query,
		acc.Username,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.EncryptedPassword,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil

}

// Function to return all cards
func (m *PostgresDBRepo) AllCards() ([]*YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Connect to db and get a list of all cards
	query := `SELECT * FROM yugioh_cards`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create new variable that is pointer to yugioh card struct
	var cards []*YugiohCard

	for rows.Next() {
		var card YugiohCard
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
func (m *PostgresDBRepo) OneCardByID(id int) ([]*YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card by id
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE id = $1`

	// Executes the query that returns rows
	rows, err := m.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create new variable which is of type slice of pointer to models.yugiohcards to hold data
	var cards []*YugiohCard

	// Create variable 'card' pointer to models.YugiohCard then iterate through rows data, Scanning into &card
	for rows.Next() {
		var card YugiohCard
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

func (m *PostgresDBRepo) GetCardByLevel(level int) ([]*YugiohCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card by id
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE level = $1`

	// Exectues the query that returns the rows
	rows, err := m.db.QueryContext(ctx, query, level)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create new variable which is of type slice of pointer to models.yugiohcard
	var cards []*YugiohCard

	// Create variable 'card' pointer to models.Yugiohcard then iterare through rows data,
	// scanning into &card

	for rows.Next() {
		var card YugiohCard
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

func (m *PostgresDBRepo) GetCardsByAttack(attack int) ([]*YugiohCard, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// Postgres query for card with attack between desired range
	query := `SELECT id, name, level, attack, defense FROM yugioh_cards WHERE attack BETWEEN $1 and $2`

	// Execute the query that returns the rows
	rows, err := m.db.QueryContext(ctx, query, attack)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Create variable which is of type slice pointer to models.yugiohcard
	var cards []*YugiohCard

	for rows.Next() {
		var card YugiohCard
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

func (m *PostgresDBRepo) AddNewCard(card YugiohCard) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `insert into yugioh_cards (name, level, attack, defense) values ($1, $2, $3, $4) returning id`

	var newID int

	err := m.db.QueryRowContext(ctx, query,
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

	_, err := m.db.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) UpdateCard(card YugiohCard) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `update yugioh_cards set name = $1, level = $2, attack = $3, defense = $4 where id = $5`

	_, err := m.db.ExecContext(ctx, query,
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

package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/uptrace/bun"
	_ "github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

// function to open DB
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pg", dsn)
	if err != nil {
		return nil, err
	}

	// verify connection to database
	err = db.Ping()

	return db, nil

}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err 
	}

	log.Println("Connected to Postgres")
	return connection, nil 
}

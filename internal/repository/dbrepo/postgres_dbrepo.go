package dbrepo

import (
	"database/sql"
	"time"
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


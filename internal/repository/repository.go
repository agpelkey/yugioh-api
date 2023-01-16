package repository

import ("database/sql")

// Database interface 
type DatabaseRepo interface {
	Connection() *sql.DB
}

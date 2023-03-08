package main

import (
	"log"
)

func main() {

	// connect to Postgres docker container
	dbConn, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	// create db tables
	if err := dbConn.Init(); err != nil {
		log.Fatal(err)
	}

	// start web server
	server := NewAPIServer(":8080", dbConn)

	// Runnit baby
	server.Run()

}

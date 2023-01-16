package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"yugioh-api/internal/repository"
	"yugioh-api/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	//for future Postgres connection
	DSN string
	DB repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// Read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=yugioh sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		 log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	log.Println("Starting application on port", port)

	// start web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

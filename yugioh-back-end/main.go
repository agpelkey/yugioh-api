package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
	DB  DatabaseRepo
}

var app application

func main() {
	// set application config
	var app application

	// Read from command line
	flag.StringVar(&app.DSN, "dsn", "postgres://postgres:yugioh@localhost:5432/yugioh?sslmode=disable", "Postgres connection string")
	flag.Parse()

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &PostgresDBRepo{DB: conn}

	defer app.DB.Connection().Close()

	log.Println("Starting application on port", port)

	// start web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
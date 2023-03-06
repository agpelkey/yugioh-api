package main

import (
	"log"
)

func main() {

	dbConn, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := dbConn.Init(); err != nil {
		log.Fatal(err)
	}

	// start web server
	server := NewAPIServer(":8080", dbConn)

	server.Run()

}

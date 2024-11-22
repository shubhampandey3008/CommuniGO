package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("There was an error creating new Database %s", err)
	}
}

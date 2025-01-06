package main

import (
	"fmt"
	"log"
	"github.com/multicurrency-tracker/backend/internal/db"
)

func main(){
	database, err := db.SetupDatabase("postgres", "lockin25", "expense_tracker", "localhost", "5432")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	fmt.Println("Hello there!")
}
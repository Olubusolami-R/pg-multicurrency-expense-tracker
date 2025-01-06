package main

import (
	"fmt"
	"log"
	"os"

	"github.com/multicurrency-tracker/backend/internal/db"
)

func main(){

	//Setup the database
	database, err := db.SetupDatabase(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "expense_tracker", "localhost", "5432")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	fmt.Println("Hello there!")
}
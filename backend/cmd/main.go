package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/db"
	"github.com/joho/godotenv"
)

func main(){

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Setup the database
	database, err := db.SetupDatabase(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()


	fmt.Println("Hello there!")
}
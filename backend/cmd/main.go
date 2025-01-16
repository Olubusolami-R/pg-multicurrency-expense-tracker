package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/apis"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/db"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Set up echo
	e := echo.New()

	e.Use(middleware.CORS())

	fmt.Println("Hello there!")

	currencyRepo := repository.NewCurrencyRepository(database)
	rateRepo:= repository.NewExchangeRateRepository(database)

	// Initialize service
	currencyService := services.NewCurrencyService(currencyRepo)
	rateService:=services.NewExchangeRateService(rateRepo)

	// Initialize handler
	currencyHandler := apis.NewCurrencyHandler(currencyService)
	rateHandler:=apis.NewExchangeRateHandler(rateService)
	
	// Optionally check and populate currencies
	err = currencyHandler.PopulateCurrencies()
	if err != nil {
		log.Fatalf("Failed to populate currencies: %v", err)
	}

	e.POST("/update-rates", rateHandler.UpdateRates)
	e.GET("/fetch-currencies", currencyHandler.GetCurrencies)
	// Start the server
	e.Logger.Fatal(e.Start(":8080"))

	
}
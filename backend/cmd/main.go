package main

import (
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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database, err := db.SetupDatabase(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	e := echo.New()
	e.Use(middleware.CORS())

	// Initialize repos
	currencyRepo := repository.NewCurrencyRepository(database)
	rateRepo:= repository.NewExchangeRateRepository(database)
	expenseRepo:=repository.NewExpenseRepository(database)

	// Initialize services
	currencyService := services.NewCurrencyService(currencyRepo)
	rateService:=services.NewExchangeRateService(rateRepo,currencyService)
	expenseService:=services.NewExpenseService(expenseRepo,currencyService)

	// Initialize handlers
	currencyHandler := apis.NewCurrencyHandler(currencyService)
	rateHandler:=apis.NewExchangeRateHandler(rateService)
	expenseHandler:=apis.NewExpenseHandler(expenseService)
	
	// Optionally check and populate currencies
	err = currencyHandler.PopulateCurrencies()
	if err != nil {
		log.Fatalf("Failed to populate currencies: %v", err)
	}

	e.POST("/update-rates", rateHandler.UpdateRates)
	e.GET("/fetch-currencies", currencyHandler.GetCurrencies)
	e.GET("/fetch-rate", rateHandler.GetExchangeRate)
	e.GET("/fetch-rates",rateHandler.GetAllExchangeRates)
	e.POST("/create-expense",expenseHandler.CreateExpense)

	e.Logger.Fatal(e.Start(":8080"))
}
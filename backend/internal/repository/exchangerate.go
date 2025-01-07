package repository

import (
	"database/sql"
)

type ExchangeRateRepository struct {
	DB *sql.DB
}

// Helper function to initialize the repository
func NewExchangeRateRepository(db *sql.DB) CurrencyRepository{
	return CurrencyRepository{DB:db}
}


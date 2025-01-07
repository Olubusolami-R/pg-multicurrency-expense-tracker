package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

type ExchangeRateRepository struct {
	DB *sql.DB
}

// Helper function to initialize the repository
func NewExchangeRateRepository(db *sql.DB) ExchangeRateRepository{
	return ExchangeRateRepository{DB:db}
}

func (r *ExchangeRateRepository) InsertSingleExchangeRate (
	baseCurrency models.Currency, 
	targetCurrency models.Currency, 
	rate float64, 
	updatedAt time.Time) error {

	query := "INSERT INTO exchange_rates (base_currency, target_currency, rate, updated_at) VALUES ($1, $2, $3, $4)"
	
	_,err := r.DB.Exec(query, baseCurrency, targetCurrency, rate, updatedAt)
	if err != nil {
		return err
	}

	return nil

}

func (r *CurrencyRepository) InsertMultipleExchangeRates (exchangeRates []models.ExchangeRate) error {

	query := "INSERT INTO exchange_rates (base_currency, target_currency, rate, updated_at) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, exchangeRate := range exchangeRates{
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*2+1, i*2+2, i*2+3, i*2+4))
		values = append(values, exchangeRate.BaseCurrency, exchangeRate.TargetCurrency, exchangeRate.Rate, exchangeRate.CreatedAt)
	}

	query += strings.Join(placeholders, ",")

	// Execute the query
	_, err := r.DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}


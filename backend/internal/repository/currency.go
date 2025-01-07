package repository

import (
	"database/sql"
	"fmt"
	"strings"

	models "github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

// Handles database operations for currency
type CurrencyRepository struct {
	DB *sql.DB
}

// Helper function to
func NewCurrencyRepository(db *sql.DB) CurrencyRepository{
	return CurrencyRepository{DB:db}
}

func (r *CurrencyRepository) InsertSingleCurrency (code string, name string) error {
	query := "INSERT INTO currencies (code, name) VALUES ($1, $2)"
	_,err := r.DB.Exec(query, code, name)
	if err != nil {
		return err
	}
	return nil

}

// Batch inserts
func (r *CurrencyRepository) InsertMultipleCurrencies (currencies []models.Currency) error {
	query := "INSERT INTO currencies (code, name) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, currency := range currencies{
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		values = append(values, currency.Code, currency.Name)
	}

	query += strings.Join(placeholders, ",")

	// Execute the query
	_, err := r.DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}
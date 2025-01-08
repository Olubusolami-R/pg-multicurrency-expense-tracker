package repository

import (
	"database/sql"
	"fmt"
	"strings"

	models "github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

type CurrencyRepository interface {
	CreateSingleCurrency(currency models.Currency) error
	CreateMultipleCurrencies(currencies []models.Currency) error
	GetCurrencies()([]models.Currency, error)
}
// Handles database operations for currency
type currencyRepository struct {
	DB *sql.DB
}

// Helper function to initialize the repository
func NewCurrencyRepository(db *sql.DB) CurrencyRepository{
	return &currencyRepository{DB:db}
}

func (r *currencyRepository) CreateSingleCurrency(currency models.Currency) error {

	query := "INSERT INTO currencies (code, name) VALUES ($1, $2)"
	_,err := r.DB.Exec(query, currency.Code, currency.Name)
	if err != nil {
		return err
	}
	return nil

}

// Batch inserts
func (r *currencyRepository) CreateMultipleCurrencies(currencies []models.Currency) error {

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

func (r *currencyRepository) GetCurrencies()([]models.Currency, error){

	// Fetch all currencies
	query:="SELECT code, name FROM currencies"

	
	rows,err:=r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currencies: %w", err)
	}
	defer rows.Close()

	var currencies []models.Currency

	//Iterating through rows
	for rows.Next() {

		var currency models.Currency

		if err := rows.Scan(&currency.Code, &currency.Name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return currencies, nil
}
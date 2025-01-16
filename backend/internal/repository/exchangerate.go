package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

type ExchangeRateRepository interface{
	CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error
	CreateMultipleExchangeRates(exchangeRates []models.ExchangeRate) error
	GetAllExchangeRates()([]models.ExchangeRate, error)
	UpsertExchangeRates(exchangeRates map[string]*models.ExchangeRate) error
	GetExchangeRate(currencyMap map[string]uint, base string, target string) (float64,error)
}

type exchangeRateRepository struct {
	DB *sql.DB
}

// Helper function to initialize the repository
func NewExchangeRateRepository(db *sql.DB) ExchangeRateRepository{
	return &exchangeRateRepository{DB:db}
}


func (r *exchangeRateRepository) CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error {

	query := "INSERT INTO exchange_rates (base_currency, target_currency, rate, updated_at) VALUES ($1, $2, $3, $4)"
	
	_,err := r.DB.Exec(query, exchangeRate.BaseCurrency, exchangeRate.TargetCurrency, exchangeRate.Rate, exchangeRate.UpdatedAt)
	if err != nil {
		return err
	}

	return nil

}

func (r *exchangeRateRepository) CreateMultipleExchangeRates(exchangeRates []models.ExchangeRate) error {

	query := "INSERT INTO exchange_rates (base_currency, target_currency, rate, updated_at) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, exchangeRate := range exchangeRates{
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*2+1, i*2+2, i*2+3, i*2+4))
		values = append(values, exchangeRate.BaseCurrency, exchangeRate.TargetCurrency, exchangeRate.Rate, exchangeRate.UpdatedAt)
	}

	query += strings.Join(placeholders, ",")

	// Execute the query
	_, err := r.DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *exchangeRateRepository) GetAllExchangeRates()([]models.ExchangeRate, error){

	query:="SELECT base_currency, target_currency, rate, updated_at FROM exchange_rates"

	rows,err:=r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currencies: %w", err)
	}
	defer rows.Close()

	var exchangeRates []models.ExchangeRate

	//Iterating through rows
	for rows.Next() {

		var exchangeRate models.ExchangeRate

		if err := rows.Scan(&exchangeRate.BaseCurrency, &exchangeRate.TargetCurrency, &exchangeRate.Rate, &exchangeRate.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		exchangeRates = append(exchangeRates, exchangeRate)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return exchangeRates, nil
}

func (r *exchangeRateRepository) UpsertExchangeRates(exchangeRates map[string]*models.ExchangeRate) error{
	// Begin transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Prepare SQL for batch insert/update
	stmt, err := tx.Prepare(`
		INSERT INTO exchange_rates (base_currency, target_currency, rate, updated_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (base_currency, target_currency) 
		DO UPDATE SET 
			rate = EXCLUDED.rate,
			updated_at = NOW();
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	for target, rate := range exchangeRates {
		_, err := stmt.Exec(rate.BaseCurrency, rate.TargetCurrency, rate.Rate)
		if err != nil {
			return fmt.Errorf("failed to insert/update exchange rate for %s: %v", target, err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	log.Println("Exchange rates updated successfully!")
	return nil

}

// Implement get single exchange rate
func (r *exchangeRateRepository) GetExchangeRate(currencyMap map[string]uint, base string, target string) (float64,error) {
	query:=`
		SELECT rate 
		FROM exchange_rates 
		WHERE base_currency = $1 AND target_currency = $2
		`
	
	rows,err:=r.DB.Query(query,currencyMap[base],currencyMap[target])
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate %s - %s: %w",base,target, err)
	}
	defer rows.Close()

	var exchangeRate float64
	for rows.Next(){
		if err := rows.Scan(&exchangeRate); err != nil {
			return 0, fmt.Errorf("failed to scan exchange rate row: %w", err)
		}
	}

	return exchangeRate,nil
}

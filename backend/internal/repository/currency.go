package repository

import "database/sql"

// Handles database operations for currency 
type CurrencyRepository struct {
	DB *sql.DB
}

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
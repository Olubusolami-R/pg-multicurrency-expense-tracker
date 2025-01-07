package repository

import (
	"database/sql"
	"time"

	models "github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

// I didn't create an interface for all the methods here. Probably should for refactoring sake and grouping sake.
// Maybe a manager.go here in repository. The interface will have same name as struct but with Cap letter starting and small letter for struct.

type ExpenseRepository struct{
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository{
	return ExpenseRepository{DB:db}
}

func (r *ExpenseRepository) InsertExpense (
	description string, 
	amount float64 , 
	currency models.Currency, 
	createdAt time.Time) error {

	query := "INSERT INTO expenses (description, amount, currency, createdAt) VALUES ($1, $2, $3, $4)"
	
	_,err := r.DB.Exec(query, description, amount, currency, createdAt)
	if err != nil {
		return err
	}

	return nil

}
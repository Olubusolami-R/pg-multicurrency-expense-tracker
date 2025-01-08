package repository

import (
	"database/sql"
	"fmt"

	models "github.com/Olubusolami-R/multicurrency-tracker/internal/models"
)

// I didn't create an interface for all the methods here. Probably should for refactoring sake and grouping sake.
// Maybe a manager.go here in repository. The interface will have same name as struct but with Cap letter starting and small letter for struct.

type ExpenseRepository interface{
	CreateExpense(expense models.Expense) error
	GetExpenses()([]models.Expense, error)
}
type expenseRepository struct{
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository{
	return &expenseRepository{DB:db}
}

func (r *expenseRepository) CreateExpense(expense models.Expense) error {

	query := "INSERT INTO expenses (description, amount, currency, createdAt) VALUES ($1, $2, $3, $4)"
	
	_,err := r.DB.Exec(query, expense.Description, expense.Amount, expense.Currency, expense.CreatedAt)
	if err != nil {
		return err
	}

	return nil

}

func (r *expenseRepository) GetExpenses()([]models.Expense, error){

	query:="SELECT description, amount, currency, createdAt FROM expenses"

	rows,err:=r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch expenses: %w", err)
	}
	defer rows.Close()

	var expenses []models.Expense

	//Iterating through rows
	for rows.Next() {

		var expense models.Expense

		if err := rows.Scan(&expense.Description, &expense.Amount, &expense.Currency, &expense.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return expenses, nil
}
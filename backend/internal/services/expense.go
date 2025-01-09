package services

import (
	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type ExpenseService interface{
	CreateExpense(expense models.Expense) error
	GetAllExpenses()([]models.Expense, error)
}

type expenseService struct{
	ExpenseRepo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) expenseService {
	return expenseService{ExpenseRepo:repo}
}

func (s *expenseService) CreateExpense(expense models.Expense) error {
	return s.ExpenseRepo.CreateExpense(expense)
} 

func (s *expenseService) GetAllExpenses()([]models.Expense, error){
	return s.ExpenseRepo.GetExpenses()
}




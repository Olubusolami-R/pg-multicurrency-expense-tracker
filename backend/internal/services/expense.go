package services

import (
	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type ExpenseService struct{
	ExpenseRepo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return ExpenseService{ExpenseRepo:repo}
}

func (s *ExpenseService) CreateExpense(expense models.Expense) error {
	return s.ExpenseRepo.CreateExpense(expense)
} 

func (s *ExpenseService) GetAllExpenses()([]models.Expense, error){
	return s.ExpenseRepo.GetExpenses()
}




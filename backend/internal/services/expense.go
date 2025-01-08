package services

import "github.com/Olubusolami-R/multicurrency-tracker/internal/repository"

type ExpenseService struct{
	ExpenseRepo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return ExpenseService{ExpenseRepo:repo}
}

func (s *ExpenseService) CreateExpense(){} // Now I have to pass a model instead




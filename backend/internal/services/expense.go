package services

import "github.com/Olubusolami-R/multicurrency-tracker/internal/repository"

type ExpenseService struct{
	repo repository.ExpenseRepository
}
func NewExpenseService (repo repository.ExpenseRepository) ExpenseService {
	return &ExpenseService{repo:repo}
}
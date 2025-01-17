package services

import (
	"fmt"
	"time"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type ExpenseService interface{
	CreateExpense(expenseData map[string]interface{}) error
	GetAllExpenses()([]models.Expense, error)
}

type expenseService struct{
	ExpenseRepo repository.ExpenseRepository
	currencyService CurrencyService
}

func NewExpenseService(repo repository.ExpenseRepository, currencyService CurrencyService) ExpenseService {
	return &expenseService{
		ExpenseRepo:repo,
		currencyService: currencyService}
}

func (s *expenseService) CreateExpense(expenseData map[string]interface{}) error {

	currencyCode:=expenseData["currency"].(string)
	
	//Because GetCurrencyIDsByCode needs a slice
	var codeSlice []string
	codeSlice = append(codeSlice, currencyCode)

	result,err:=s.currencyService.GetCurrencyIDsByCode(codeSlice)
	if err!=nil{
		fmt.Println("Error getting expense currency code")
		return fmt.Errorf("error getting expense currency code: %w",err)
	}

	// now create expense object
	expenseCurrencyID:= result[currencyCode]

	expense:=&models.Expense{
		Description: expenseData["description"].(string),
		Amount: expenseData["amount"].(float64),
		Currency: expenseCurrencyID,
		CreatedAt: time.Now()}
	
	return s.ExpenseRepo.CreateExpense(expense)
} 

func (s *expenseService) GetAllExpenses()([]models.Expense, error){
	return s.ExpenseRepo.GetExpenses()
}




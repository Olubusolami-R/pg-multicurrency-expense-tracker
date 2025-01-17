package apis

import (
	"fmt"
	"net/http"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler interface {
	CreateExpense(c echo.Context) error
	GetAllExpenses(c echo.Context) error
}

type expenseHandler struct {
	expenseService services.ExpenseService
}

func NewExpenseHandler(service services.ExpenseService) ExpenseHandler {
	return &expenseHandler{expenseService: service}
}

func (h *expenseHandler) CreateExpense(c echo.Context) error {

	var requestData struct {
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
		Currency    string    `json:"currency"`
	}

	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body."})
	}

	requestMap:=make(map[string]interface{})

	requestMap["description"]=requestData.Description
	requestMap["amount"]=requestData.Amount
	requestMap["currency"]=requestData.Currency

	err:=h.expenseService.CreateExpense(requestMap)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,fmt.Errorf("error creating expense: %w",err))
	}

	fmt.Println("Expense created!")

	return c.JSON(http.StatusOK,requestMap)
}

func (h *expenseHandler) GetAllExpenses(c echo.Context) error{
	expenses,err:=h.expenseService.GetAllExpenses()
	if err!=nil{
		fmt.Println(err)
		return fmt.Errorf("error getting all expenses: %w", err)
	}
	return c.JSON(http.StatusOK,expenses)
}


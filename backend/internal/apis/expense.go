package apis

import (
	"fmt"
	"net/http"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler interface {
	CreateExpense(c echo.Context) error
}

type expenseHandler struct {
	expenseService services.ExpenseService
}

func NewExpenseHandler(service services.ExpenseService) ExpenseHandler {
	return &expenseHandler{expenseService: service}
}

func (h *expenseHandler) CreateExpense(c echo.Context) error {
	fmt.Println("debug 1")
	var requestData struct {
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
		Currency    string    `json:"currency"`
	}

	if err := c.Bind(&requestData); err != nil {
		fmt.Println("debug 2: error here on failing to bind")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body."})
	}

	requestMap:=make(map[string]interface{})

	requestMap["description"]=requestData.Description
	requestMap["amount"]=requestData.Amount
	requestMap["currency"]=requestData.Currency

	fmt.Println("debug 3:",requestMap)

	err:=h.expenseService.CreateExpense(requestMap)
	if err!=nil{
		fmt.Println("debug x: error here on creating expense")
		return c.JSON(http.StatusInternalServerError,fmt.Errorf("error creating expense: %w",err))
	}

	fmt.Println("Expense created!")

	return c.JSON(http.StatusOK,requestMap)
}

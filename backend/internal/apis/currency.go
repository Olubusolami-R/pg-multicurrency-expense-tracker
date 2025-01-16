package apis

import (
	"fmt"
	"net/http"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
	"github.com/labstack/echo/v4"
)

type CurrencyHandler interface{
	PopulateCurrencies() error
	GetCurrencies(c echo.Context)error
}

type currencyHandler struct {
	currencyService services.CurrencyService
}

func NewCurrencyHandler(service services.CurrencyService) CurrencyHandler {
	return &currencyHandler{currencyService: service}
}
// called once
func (h *currencyHandler) PopulateCurrencies() error {
	
	alreadyPopulated, err := h.currencyService.CheckCurrenciesPopulated() // e.g., count rows in DB
    if err != nil {
        return fmt.Errorf("failed to check currencies: %w", err)
    }

    if alreadyPopulated {
        fmt.Println("Currencies are already populated, skipping.")
        return nil
    }
	
	//If not populated, then populate
	currencies,err:=h.currencyService.LoadCurrencies()
	if err!=nil{
		return fmt.Errorf("error loading currencies :%w", err)
	}

	err=h.currencyService.CreateMultipleCurrencies(currencies)
	if err!=nil{
		return fmt.Errorf("error populating currencies table:%w", err)
	}

	fmt.Println("Currencies table successfully populated! Check psql.")
	return nil
}


func (h *currencyHandler) GetCurrencies(c echo.Context)error{
	currencies,err:= h.currencyService.GetAllCurrencies()
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,fmt.Errorf("error fetching all currencies :%w", err))
	}
	fmt.Println(currencies)
	return c.JSON(http.StatusOK, "Currencies fetched, Check terminal.")
}
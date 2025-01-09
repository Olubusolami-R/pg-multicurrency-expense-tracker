package apis

import (
	"fmt"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
)

type CurrencyHandler interface{

}

type currencyHandler struct {
	currencyService services.CurrencyService
}

// called once
func (h *currencyHandler) PopulateCurrencies() error {
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

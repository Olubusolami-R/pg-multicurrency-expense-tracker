package apis

import (
	"fmt"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
)

type CurrencyHandler interface{
	PopulateCurrencies() error
}

type currencyHandler struct {
	currencyService services.CurrencyService
}

// called once
func (h *currencyHandler) PopulateCurrencies() error {

	//have to check if populated already so needs service and repo haewww

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






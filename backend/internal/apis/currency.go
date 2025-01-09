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






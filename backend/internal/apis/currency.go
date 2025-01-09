package apis

import "github.com/Olubusolami-R/multicurrency-tracker/internal/services"

type CurrencyHandler interface{

}

type currencyHandler struct {
	currencyService services.CurrencyService
}

// called once
func (h *currencyHandler) ProcessCurrencies() error {

}

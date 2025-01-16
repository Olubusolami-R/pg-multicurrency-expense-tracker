package apis

import (
	"fmt"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
)

type ExchangeRateHandler interface{
	PopulateCurrencies() error
}

type exchangeRateHandler struct {
	exchangeRateService services.ExchangeRateService
}

func NewExchangeRateHandler(service services.CurrencyService) CurrencyHandler {
	return &currencyHandler{currencyService: service}
}

func (h *exchangeRateHandler) UpdateRates()error{
	ratesJSON,err:=h.exchangeRateService.CallExchangeRateAPI()
	if err!=nil{
		return fmt.Errorf("error fetching latest exchange rates from API")
	}

	ratesMap,err:=h.exchangeRateService.ProcessAPIOutput(ratesJSON)
	if err!=nil{
		return fmt.Errorf("error processing exchange rate API output into map")
	}

	err=h.exchangeRateService.UpsertExchangeRates(ratesMap)
	if err!=nil{
		return fmt.Errorf("error upserting exchange rates in database")
	}

	fmt.Println("Rates updated successfully! Check Postgres")

	return nil
}
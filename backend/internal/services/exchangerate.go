package services

import (
	"encoding/json"
	"fmt"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type ExchangeRateService struct{
	Repo repository.ExchangeRateRepository
	currencyService CurrencyService
}

func NewExchangeRateService(repo repository.ExchangeRateRepository)ExchangeRateService{
	return ExchangeRateService{Repo: repo}
}

func (s *ExchangeRateService) ProcessAPIOutput(jsonData []byte)([]models.ExchangeRate,error){
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil,err
	}

	// Retrieve the base currency
	baseCurrency:=data["base"].(string)

	// Convert to list format and retrieve code-ID map of base Currency
	var baseCurrencyList []string // Use []string directly
	baseCurrencyList = append(baseCurrencyList, baseCurrency)
	baseCurrencyMap, err:=s.currencyService.GetCurrencyIDsByCode(baseCurrencyList)
	if err != nil {
		return nil, fmt.Errorf("error fetching base currency ID: %w", err)
	}

	// retrieve the targetCurrencyCode-rate map
	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("rates field missing or invalid")
	}

	var codes []string
	for code := range rates {
		codes = append(codes, code)
	}

	// Now focus on retrieving code-ID map for all target currencies
	currencyMap, err := s.currencyService.GetCurrencyIDsByCode(codes)
	if err != nil {
		return nil, fmt.Errorf("error fetching currencies: %w", err)
	}

	// Convert currencyMap items to storable exchangeRate formats.
	var exchangeRates []models.ExchangeRate
	for code, rate :=range rates {
		var exchangeRate models.ExchangeRate

		// Assert rate as float64
		rateValue, ok := rate.(float64)
		if !ok {
			return nil, fmt.Errorf("rate for code %s is not a valid float64", code)
		}

		exchangeRate.BaseCurrency=*baseCurrencyMap[baseCurrency]
		exchangeRate.TargetCurrency=*currencyMap[code]
		exchangeRate.Rate=rateValue

		exchangeRates=append(exchangeRates, exchangeRate)
	}

	return exchangeRates,nil
}

func (s *ExchangeRateService) CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error {
	return s.Repo.CreateSingleExchangeRate(exchangeRate)
}

func (s *ExchangeRateService) CreateMultipleExchangeRates(exchangeRates []models.ExchangeRate) error {
	return s.Repo.CreateMultipleExchangeRates(exchangeRates)
}

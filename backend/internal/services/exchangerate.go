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
	baseCurrency:=data["base"]

	//Convert to list format 
	var baseCurrencyList []interface{}
	baseCurrencyList=append(baseCurrencyList, baseCurrency)
	
	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("rates field missing or invalid")
	}

	var codes []string
	for code := range rates {
		codes = append(codes, code)
	}

	currencyMap, err := s.currencyService.GetCurrenciesByCode(codes)
	if err != nil {
		return nil, fmt.Errorf("error fetching currencies: %w", err)
	}

	// Convert currencyMap items to storable exchangeRate formats.


}

func (s *ExchangeRateService) CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error {
	return s.Repo.CreateSingleExchangeRate(exchangeRate)
}

func (s *ExchangeRateService) CreateMultipleExchangeRates(exchangeRates []models.ExchangeRate) error {
	return s.Repo.CreateMultipleExchangeRates(exchangeRates)
}

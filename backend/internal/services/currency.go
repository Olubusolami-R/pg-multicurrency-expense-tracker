package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type CurrencyService struct {
	CurrencyRepo repository.CurrencyRepository
}

func LoadCurrencies() ([]models.Currency, error) {
	path := filepath.Join("internal", "resources", "currencies.json")

	// Read the file content
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("could not read currencies.json: %v", err)
    }

	// Parse JSON data into a map
    var currencyMap map[string]string
    err = json.Unmarshal(data, &currencyMap)
    if err != nil {
        return nil, fmt.Errorf("could not parse currencies.json: %v", err)
    }

    // Convert map to a slice of Currency structs
    var currencies []models.Currency
    for code, name := range currencyMap {
        currency := models.Currency{
            Code: code,
            Name: name,
        }
        currencies = append(currencies, currency)
    }

    return currencies, nil
}

func NewCurrencyService(repo repository.CurrencyRepository) CurrencyService {
	return CurrencyService{CurrencyRepo: repo}
}

func (s *CurrencyService) CreateSingleCurrency(currency models.Currency) error {
	return s.CurrencyRepo.CreateSingleCurrency(currency)
}

func (s *CurrencyService) CreateMultipleCurrencies(currencies []models.Currency) error {
	return s.CurrencyRepo.CreateMultipleCurrencies(currencies)
}

func (s *CurrencyService) GetAllCurrencies()([]models.Currency, error){
	return s.CurrencyRepo.GetCurrencies()
}

func (s *CurrencyService) GetCurrencyIDsByCode(codes []string)(map[string]*uint, error){
	return s.CurrencyRepo.GetCurrencyIDsByCode(codes)
}
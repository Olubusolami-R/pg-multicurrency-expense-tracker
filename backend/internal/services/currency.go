package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type CurrencyService interface{
	LoadCurrencies() ([]models.Currency, error)
	CreateSingleCurrency(currency models.Currency) error
	CreateMultipleCurrencies(currencies []models.Currency) error
	GetAllCurrencies()([]models.Currency, error)
	GetCurrencyIDsByCode(codes []string)(map[string]uint, error)
	CheckCurrenciesPopulated() (bool,error)
}

type currencyService struct {
	CurrencyRepo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) CurrencyService {
	return &currencyService{CurrencyRepo: repo}
}

func (s *currencyService) CheckCurrenciesPopulated() (bool,error) {
	return s.CurrencyRepo.CheckCurrenciesPopulated()
}

func (s *currencyService) LoadCurrencies() ([]models.Currency, error) {
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

func (s *currencyService) CreateSingleCurrency(currency models.Currency) error {
	return s.CurrencyRepo.CreateSingleCurrency(currency)
}

func (s *currencyService) CreateMultipleCurrencies(currencies []models.Currency) error {
	return s.CurrencyRepo.CreateMultipleCurrencies(currencies)
}

func (s *currencyService) GetAllCurrencies()([]models.Currency, error){
	fmt.Printf("currencyService: %+v\n", s)
	if s.CurrencyRepo == nil {
		fmt.Println("currency repository is not initialized")
		return nil, fmt.Errorf("currency repository is not initialized")
	}
	return s.CurrencyRepo.GetCurrencies()
}

//Problem location - why is currencyRepo returning nil???
func (s *currencyService) GetCurrencyIDsByCode(codes []string)(map[string]uint, error){
	fmt.Println("inside GetCurrency service step 3")
	fmt.Printf("currencyService: %+v\n", s)
	fmt.Println("This is codes:",codes)
	if s.CurrencyRepo == nil {
		fmt.Println("currency repository is not initialized")
		return nil, fmt.Errorf("currency repository is not initialized")
	}
	return s.CurrencyRepo.GetCurrencyIDsByCode(codes)
}
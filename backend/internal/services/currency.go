package services

import (
	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type CurrencyService struct {
	CurrencyRepo repository.CurrencyRepository
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

func (s *CurrencyService) GetCurrenciesBySymbols(codes []string)([]models.Currency, error){
	return s.CurrencyRepo.GetCurrenciesByCode(codes)
}
package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/Olubusolami-R/multicurrency-tracker/internal/repository"
)

type ExchangeRateService interface{
	CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error
	CallExchangeRateAPI() ([]byte,error)
	ProcessAPIOutput(jsonData []byte)(map[string]*models.ExchangeRate,error)
	UpsertExchangeRates(exchangeRates map[string]*models.ExchangeRate) error
	GetExchangeRate(base string, target string)(float64, error)
	GetAllExchangeRates()([]models.ExchangeRate, error)
}

type exchangeRateService struct{
	Repo repository.ExchangeRateRepository
	currencyService CurrencyService
}

func NewExchangeRateService(repo repository.ExchangeRateRepository, currencyService CurrencyService)ExchangeRateService{
	return &exchangeRateService{
		Repo: repo,
		currencyService: currencyService}
}

func (s *exchangeRateService) CallExchangeRateAPI() ([]byte,error) {

	accessKey := os.Getenv("EXCHANGE_API_KEY")
	if accessKey == "" {
		fmt.Println("Error: EXCHANGE_API_KEY environment variable is not set")
		return nil,nil
	}
	
	apiURL:=fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s",os.Getenv("EXCHANGE_API_KEY"))
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil,fmt.Errorf("failed to fetch exchange rates: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil

}

func (s *exchangeRateService) ProcessAPIOutput(jsonData []byte)(map[string]*models.ExchangeRate,error){
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil,err
	}

	baseCurrency:=data["base"].(string)

	// Convert to list format and retrieve code-ID map of base Currency
	var baseCurrencyList []string 
	baseCurrencyList = append(baseCurrencyList, baseCurrency)

	if len(baseCurrencyList) == 0 {
		return nil, fmt.Errorf("no baseCurrency provided")
	}
	
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
	exchangeRates:= make(map[string]*models.ExchangeRate)
	for code, rate :=range rates {
		var exchangeRate models.ExchangeRate

		rateValue, ok := rate.(float64)
		if !ok {
			return nil, fmt.Errorf("rate for code %s is not a valid float64", code)
		}

		exchangeRate.BaseCurrency=baseCurrencyMap[baseCurrency]
		exchangeRate.TargetCurrency=currencyMap[code]
		exchangeRate.Rate=rateValue

		exchangeRates[code]=&exchangeRate
	}

	return exchangeRates,nil
}

func (s *exchangeRateService) CreateSingleExchangeRate(exchangeRate models.ExchangeRate) error {
	return s.Repo.CreateSingleExchangeRate(exchangeRate) //handler not created
}

func (s *exchangeRateService) UpsertExchangeRates(exchangeRates map[string]*models.ExchangeRate) error{
	return s.Repo.UpsertExchangeRates(exchangeRates)
}

func (s *exchangeRateService) GetExchangeRate(base string, target string)(float64, error){
	var currencies []string
	currencies = append(currencies, base)
	currencies = append(currencies, target)

	currencyMap,err:=s.currencyService.GetCurrencyIDsByCode(currencies)
	if err!=nil{
		return -1,fmt.Errorf("unable to retrieve the IDs of the base and target currency: %w", err)
	}

	return s.Repo.GetExchangeRate(currencyMap, base, target)
}

func (s *exchangeRateService) GetAllExchangeRates()([]models.ExchangeRate, error){
	return s.Repo.GetAllExchangeRates()
}
package apis

import (
	"fmt"
	"net/http"

	"github.com/Olubusolami-R/multicurrency-tracker/internal/services"
	"github.com/labstack/echo/v4"
)

type ExchangeRateHandler interface{
	UpdateRates(c echo.Context)error
}

type exchangeRateHandler struct {
	exchangeRateService services.ExchangeRateService
}

func NewExchangeRateHandler(service services.ExchangeRateService) ExchangeRateHandler {
	return &exchangeRateHandler{exchangeRateService: service}
}

func (h *exchangeRateHandler) UpdateRates(c echo.Context)error{
	ratesJSON,err:=h.exchangeRateService.CallExchangeRateAPI()
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("error fetching latest exchange rates from API"))
	}

	ratesMap,err:=h.exchangeRateService.ProcessAPIOutput(ratesJSON)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,fmt.Errorf("error processing exchange rate API output into map"))
	}

	err=h.exchangeRateService.UpsertExchangeRates(ratesMap)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,fmt.Errorf("error upserting exchange rates in database"))
	}

	fmt.Println("Rates updated successfully! Check Postgres")

	return c.JSON(http.StatusOK, "Rates updated successfully! Check Postgres.")
}
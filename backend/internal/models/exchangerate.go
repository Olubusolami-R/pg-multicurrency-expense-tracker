package models

import "time"

type ExchangeRate struct {
	ID             uint      `json:"id"`
	BaseCurrency   uint      `json:"base_currency"`
	TargetCurrency uint      `json:"target_currency"`
	Rate           float64   `json:"rate"`
	UpdatedAt      time.Time `json:"updated_at"`
}

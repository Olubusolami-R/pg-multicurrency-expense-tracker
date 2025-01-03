package models

import "time"

type Expense struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Currency    uint      `json:"currency"`
	CreatedAt   time.Time `json:"created_at"`
}

package models

// Step 1 in setting up the backend

type Currency struct{
	ID   uint   `json:"id"`
	Code string `json:"code"` // Currency code (e.g., USD, EUR)
	Name string `json:"name"`
}
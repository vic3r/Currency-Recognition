package models

// Expense is a job structure defined
type Expense struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	CurrencyType string `json:"currency_type"`
	Total        int    `json:"total"`
}

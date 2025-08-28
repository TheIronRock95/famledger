package models

import "time"

// Income entry
type Income struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Source    string    `json:"source"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// Expense entry
type Expense struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Category  string    `json:"category"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// Goal entry
type Goal struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Name          string    `json:"name"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	TargetDate    time.Time `json:"target_date"`
	CreatedAt     time.Time `json:"created_at"`
}

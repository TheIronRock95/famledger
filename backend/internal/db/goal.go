package db

import "time"

type Goal struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Name          string    `json:"name"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	TargetDate    time.Time `json:"target_date"`
	CreatedAt     time.Time `json:"created_at"`
}

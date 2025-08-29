package db

import "time"

type Income struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Source    string    `json:"source"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

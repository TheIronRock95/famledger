package api

import (
	"encoding/json"
	"net/http"

	"github.com/TheIronRock95/famledger/internal/db"
)

// AddIncome adds a new income entry, letting the DB generate id and created_at
func AddIncome(w http.ResponseWriter, r *http.Request) {
	var inc struct {
		UserID string  `json:"user_id"`
		Source string  `json:"source"`
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert only the fields that exist in the table
	var result interface{}
	err := db.Supabase.DB.From("income").Insert(map[string]interface{}{
		"user_id": inc.UserID,
		"source":  inc.Source,
		"amount":  inc.Amount,
	}).Execute(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetIncome retrieves all income entries
func GetIncome(w http.ResponseWriter, r *http.Request) {
	var income []map[string]interface{}

	err := db.Supabase.DB.From("income").Select("*").Execute(&income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(income)
}

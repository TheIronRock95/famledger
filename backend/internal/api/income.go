package api

import (
	"encoding/json"
	"net/http"

	"github.com/TheIronRock95/famledger/internal/db"
)

// AddIncome handler
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

	result, err := db.AddIncomeRecord(inc.UserID, inc.Source, inc.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// GetIncome handler
func GetIncome(w http.ResponseWriter, r *http.Request) {
	income, err := db.GetIncomeRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(income)
}

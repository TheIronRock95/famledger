package api

import (
	"encoding/json"
	"net/http"

	"github.com/TheIronRock95/famledger/internal/db"
	"github.com/TheIronRock95/famledger/internal/models"
)

// AddExpense handles POST /expenses
func AddExpense(w http.ResponseWriter, r *http.Request) {
	var exp models.Expense
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.InsertExpense(&exp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exp)
}

// GetExpenses handles GET /expenses
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := db.GetAllExpenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}

package api

import (
	"encoding/json"
	"net/http"

	"github.com/TheIronRock95/famledger/backend/internal/db"
	"github.com/TheIronRock95/famledger/backend/internal/models"
)

// Add income
func AddIncome(w http.ResponseWriter, r *http.Request) {
	var inc models.Income
	if err := json.NewDecoder(r.Body).Decode(&inc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Supabase.DB.From("income").Insert(inc).Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inc)
}

// Get all income
func GetIncome(w http.ResponseWriter, r *http.Request) {
	var income []models.Income

	err := db.Supabase.DB.From("income").Select("*").Execute(&income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(income)
}

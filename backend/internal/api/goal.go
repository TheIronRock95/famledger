package api

import (
	"encoding/json"
	"net/http"

	"github.com/TheIronRock95/famledger/internal/db"
)

// AddGoal endpoint
func AddGoal(w http.ResponseWriter, r *http.Request) {
	var g db.Goal
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.InsertGoal(&g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(g)
}

// GetGoals endpoint
func GetGoals(w http.ResponseWriter, r *http.Request) {
	goals, err := db.GetAllGoals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(goals)
}

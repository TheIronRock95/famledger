package api

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	// Income endpoints
	r.HandleFunc("/income", AddIncome).Methods("POST")
	r.HandleFunc("/income", GetIncome).Methods("GET")

	// Expenses endpoints
	r.HandleFunc("/expenses", AddExpense).Methods("POST")
	r.HandleFunc("/expenses", GetExpenses).Methods("GET")

	// Goals endpoints
	r.HandleFunc("/goal", AddGoal).Methods("POST")
	r.HandleFunc("/goals", GetGoals).Methods("GET")
}

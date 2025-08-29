package db

import (
	"log"

	"github.com/TheIronRock95/famledger/internal/models"
)

// InsertExpense inserts a new expense into Supabase
func InsertExpense(expense *models.Expense) error {
	var result interface{}
	err := Supabase.DB.From("expenses").Insert(map[string]interface{}{
		"user_id":  expense.UserID,
		"category": expense.Category,
		"amount":   expense.Amount,
	}).Execute(&result)

	if err != nil {
		log.Printf("Error inserting expense: %v", err)
		return err
	}
	return nil
}

// GetAllExpenses fetches all expenses from Supabase
func GetAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	err := Supabase.DB.From("expenses").Select("*").Execute(&expenses)
	if err != nil {
		log.Printf("Error fetching expenses: %v", err)
		return nil, err
	}
	return expenses, nil
}

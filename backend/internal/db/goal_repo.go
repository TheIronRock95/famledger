package db

import (
	"log"
)

// InsertGoal inserts a new goal into Supabase
func InsertGoal(goal *Goal) error {
	var result interface{}
	err := Supabase.DB.From("goals").
		Insert(map[string]interface{}{
			"user_id":        goal.UserID,
			"name":           goal.Name,
			"target_amount":  goal.TargetAmount,
			"current_amount": goal.CurrentAmount,
			"target_date":    goal.TargetDate,
		}).Execute(&result)

	if err != nil {
		log.Printf("Error inserting goal: %v", err)
		return err
	}
	return nil
}

// GetAllGoals fetches all goals from Supabase
func GetAllGoals() ([]Goal, error) {
	var goals []Goal
	err := Supabase.DB.From("goals").Select("*").Execute(&goals)
	if err != nil {
		log.Printf("Error fetching goals: %v", err)
		return nil, err
	}
	return goals, nil
}

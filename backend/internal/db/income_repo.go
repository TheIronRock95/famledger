package db

// AddIncome inserts a new income record into Supabase
func AddIncomeRecord(userID, source string, amount float64) (*Income, error) {
	var result []Income

	err := Supabase.DB.From("income").
		Insert(map[string]interface{}{
			"user_id": userID,
			"source":  source,
			"amount":  amount,
		}).
		Execute(&result)

	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, nil
}

// GetIncomeRecords fetches all income entries
func GetIncomeRecords() ([]Income, error) {
	var income []Income

	err := Supabase.DB.From("income").
		Select("*").
		Execute(&income)

	if err != nil {
		return nil, err
	}

	return income, nil
}

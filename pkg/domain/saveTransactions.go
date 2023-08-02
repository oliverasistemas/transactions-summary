package domain

import (
	"fmt"
	"stori/pkg/data"
)

// saveTransactions saves transactions to MySQL database.
func saveTransactions(transactions []Transaction) error {
	db, err := data.ConnectMySQL()
	if err != nil {
		return fmt.Errorf("error getting connection: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO transactions (transaction_date, transaction_amount, transaction_type, account_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	for _, transaction := range transactions {
		_, err = stmt.Exec(transaction.Date, transaction.Amount, transaction.Type, transaction.AccountID)
		if err != nil {
			return fmt.Errorf("error executing query: %w", err)
		}
	}

	return nil
}

package domain

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func processReader(reader *csv.Reader) ([]Transaction, error) {
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	var transactions []Transaction

	for {
		var transactionType string

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		date, err := parseDate(record)
		if err != nil {
			return nil, err
		}
		amount, err := strconv.ParseFloat(strings.TrimPrefix(record[2], "+"), 64)
		if err != nil {
			return nil, err
		}
		accountID, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(record[2], "+") {
			transactionType = TransactionTypeCredit
		} else {
			transactionType = TransactionTypeDebit
		}

		transaction := Transaction{
			ID:        id,
			Date:      date,
			Amount:    amount,
			Type:      transactionType,
			AccountID: uint(accountID),
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func parseDate(record []string) (time.Time, error) {
	dateFormats := []string{"1/2/2006", "1/2"}

	for _, format := range dateFormats {
		date, err := time.Parse(format, record[1])
		if err == nil {
			// Check if year is not specified
			if date.Year() == 0 {
				// Update year to the current year
				currentYear := time.Now().Year()
				date = date.AddDate(currentYear, 0, 0)
			}
			return date, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", record[1])
}

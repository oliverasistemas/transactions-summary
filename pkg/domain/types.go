package domain

import (
	"time"
)

const (
	TransactionTypeCredit = "credit"
	TransactionTypeDebit  = "debit"
)

// Account represents a Stori account
type Account struct {
	ID   int
	Name string
}

// Transaction represents a transaction record
type Transaction struct {
	ID        int
	Date      time.Time
	Amount    float64
	Type      string
	AccountID uint
}

// Summary struct to store the summary information
type Summary struct {
	AvgCreditAmount float64
	AvgDebitAmount  float64
	Months          map[string]int
	NumCredit       int
	NumDebit        int
	TotalBalance    float64
}

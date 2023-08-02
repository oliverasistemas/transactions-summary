package domain

import (
	"testing"
	"time"
)

func TestCalculateSummary(t *testing.T) {
	transactions := []Transaction{
		{ID: 1, Date: time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC), Amount: 1000.0, Type: TransactionTypeCredit},
		{ID: 2, Date: time.Date(2023, time.February, 20, 0, 0, 0, 0, time.UTC), Amount: 500.0, Type: TransactionTypeCredit},
		{ID: 3, Date: time.Date(2023, time.January, 10, 0, 0, 0, 0, time.UTC), Amount: 300.0, Type: TransactionTypeDebit},
		{ID: 4, Date: time.Date(2023, time.February, 5, 0, 0, 0, 0, time.UTC), Amount: 200.0, Type: TransactionTypeDebit},
		{ID: 5, Date: time.Date(2023, time.March, 25, 0, 0, 0, 0, time.UTC), Amount: 800.0, Type: TransactionTypeCredit},
	}

	summary := calculateSummary(transactions)

	// Test total balance
	expectedTotalBalance := 2800.0
	if summary.TotalBalance != expectedTotalBalance {
		t.Errorf("Expected TotalBalance to be %.2f, but got %.2f", expectedTotalBalance, summary.TotalBalance)
	}

	// Test average credit amount
	expectedAvgCreditAmount := (1000.0 + 500.0 + 800.0) / 3.0
	if summary.AvgCreditAmount != expectedAvgCreditAmount {
		t.Errorf("Expected AvgCreditAmount to be %.2f, but got %.2f", expectedAvgCreditAmount, summary.AvgCreditAmount)
	}

	// Test average debit amount
	expectedAvgDebitAmount := (300.0 + 200.0) / 2.0
	if summary.AvgDebitAmount != expectedAvgDebitAmount {
		t.Errorf("Expected AvgDebitAmount to be %.2f, but got %.2f", expectedAvgDebitAmount, summary.AvgDebitAmount)
	}

	// Test number of credit transactions
	expectedNumCredit := 3
	if summary.NumCredit != expectedNumCredit {
		t.Errorf("Expected NumCredit to be %d, but got %d", expectedNumCredit, summary.NumCredit)
	}

	// Test number of debit transactions
	expectedNumDebit := 2
	if summary.NumDebit != expectedNumDebit {
		t.Errorf("Expected NumDebit to be %d, but got %d", expectedNumDebit, summary.NumDebit)
	}

	// Test months map
	expectedMonths := map[string]int{
		"January 2023":  2,
		"February 2023": 2,
		"March 2023":    1,
	}
	for month, count := range expectedMonths {
		if summary.Months[month] != count {
			t.Errorf("Expected %s count to be %d, but got %d", month, count, summary.Months[month])
		}
	}
}

package domain

import (
	"os"
	"testing"
	"time"
)

func TestExtractTransactions_LocalFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testcsv")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tmpFile.Close()

	testData := `ID,Date,Amount,AccountID
1,1/2/2023,+100,1
2,3/4/2023,-10,2`
	_, err = tmpFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Error writing test data to temporary file: %v", err)
	}
	localFile := tmpFile.Name()
	defer os.Remove(localFile)

	os.Setenv(localCSVFilePath, localFile)
	defer os.Unsetenv(localCSVFilePath)

	transactions, err := extractTransactions()
	if err != nil {
		t.Fatalf("Error extracting transactions from local file: %v", err)
	}

	expected := []Transaction{
		{ID: 1, Date: time.Date(time.Now().Year(), 1, 2, 0, 0, 0, 0, time.UTC), Amount: 100, AccountID: 1, Type: TransactionTypeCredit},
		{ID: 2, Date: time.Date(time.Now().Year(), 3, 4, 0, 0, 0, 0, time.UTC), Amount: -10, AccountID: 2, Type: TransactionTypeDebit},
	}
	if len(transactions) != len(expected) {
		t.Fatalf("Expected %d transactions, but got %d", len(expected), len(transactions))
	}
	for i, txn := range transactions {
		if txn != expected[i] {
			t.Errorf("Expected transaction %v, but got %v", expected[i], txn)
		}
	}
}

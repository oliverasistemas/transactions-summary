package domain

import (
	"encoding/csv"
	"reflect"
	"strings"
	"testing"
	"time"
)

// Helper function to create a CSV reader from a string
func createCSVReader(data string) *csv.Reader {
	return csv.NewReader(strings.NewReader(data))
}

func TestProcessReader_ValidInput(t *testing.T) {
	csvData := `ID,Date,Amount,AccountID
1,1/2/2023,100,1
2,3/4/2023,200.50,1`

	reader := createCSVReader(csvData)

	expectedTransactions := []Transaction{
		{ID: 1, Date: time.Date(time.Now().Year(), time.January, 2, 0, 0, 0, 0, time.UTC), Amount: 100, Type: TransactionTypeDebit, AccountID: 1},
		{ID: 2, Date: time.Date(time.Now().Year(), time.March, 4, 0, 0, 0, 0, time.UTC), Amount: 200.50, Type: TransactionTypeDebit, AccountID: 1},
	}

	transactions, err := processReader(reader)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if !reflect.DeepEqual(transactions, expectedTransactions) {
		t.Errorf("Expected transactions:\n%v\nBut got transactions:\n%v", expectedTransactions, transactions)
	}
}

func TestProcessReader_InvalidCSV(t *testing.T) {
	csvData := `ID,Date,Amount,AccountID
1,1/2,100,1
2,3/4,1`

	reader := createCSVReader(csvData)

	_, err := processReader(reader)
	if err == nil {
		t.Error("Expected an error for invalid CSV, but got none.")
	}
}

func TestProcessReader_EmptyInput(t *testing.T) {
	csvData := `ID,Date,Amount,AccountID`

	reader := createCSVReader(csvData)

	transactions, err := processReader(reader)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if len(transactions) != 0 {
		t.Errorf("Expected empty transactions slice, but got: %v", transactions)
	}
}

func TestProcessReader_InvalidDataTypes(t *testing.T) {
	csvData := `ID,Date,Amount,AccountID
1,1/2,abc,1
2,3/4,200.50,1`

	reader := createCSVReader(csvData)

	_, err := processReader(reader)
	if err == nil {
		t.Error("Expected an error for invalid data types, but got none.")
	}
}

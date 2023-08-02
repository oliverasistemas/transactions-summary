package domain

import (
	"fmt"
	"regexp"
	"stori/pkg/mailer"
)

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailPattern)
	return regex.MatchString(email)
}

// SendSummary reads and processes transactions from an S3 file,
// saves them to the MySQL database, and sends an email with a summary of the transactions.
// Returns an error if any step fails.
func SendSummary(address string) error {

	if !isValidEmail(address) {
		return fmt.Errorf("invalid email address: %s", address)
	}

	transactions, err := extractTransactions()
	if err != nil {
		return fmt.Errorf("error reading and processing file from S3: %w", err)
	}

	err = saveTransactions(transactions)
	if err != nil {
		return fmt.Errorf("error saving data to db: %w", err)
	}

	return sendSummaryEmail(transactions, address)
}

func sendSummaryEmail(transactions []Transaction, address string) error {
	return mailer.NewMailer().Send(getMailContent(calculateSummary(transactions)), address)
}

package domain

// calculateSummary calculates a summary of financial transactions based on the provided slice of Transaction objects.
// It computes various statistics such as total balance, number of credit and debit transactions,
// average credit and debit amounts, and the count of transactions in each month.
// The function returns a Summary object that holds the computed summary data.
func calculateSummary(transactions []Transaction) Summary {
	summary := Summary{}
	summary.Months = make(map[string]int)

	for _, transaction := range transactions {
		summary.TotalBalance += transaction.Amount
		if transaction.Type == TransactionTypeCredit {
			summary.NumCredit++
			summary.AvgCreditAmount += transaction.Amount
		} else {
			summary.NumDebit++
			summary.AvgDebitAmount += transaction.Amount
		}

		monthKey := transaction.Date.Format("January 2006")
		summary.Months[monthKey]++
	}

	// Calculate average amounts
	if summary.NumCredit > 0 {
		summary.AvgCreditAmount /= float64(summary.NumCredit)
	}
	if summary.NumDebit > 0 {
		summary.AvgDebitAmount /= float64(summary.NumDebit)
	}

	return summary
}

package main

import (
	"flag"
	"fmt"
	"os"
	"stori/pkg/domain"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "", "Summary recipients email address")

	// Parse command-line flags
	flag.Parse()

	// Check if the mandatory flag is provided
	if address == "" {
		fmt.Println("Error: The mandatory flag 'address' is missing.")
		flag.Usage()
		os.Exit(1)
	}

	err := domain.SendSummary(address)
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Summary email sent successfully.")
	}
}

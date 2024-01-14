package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Initialize a scanner to read input from Stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Print the header for the output
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// Infinite loop for continuous input until Ctrl+C is pressed
	for {
		// Prompt the user for input
		fmt.Print("Enter domain (or press Ctrl+C to exit): ")

		// Read the next line of input
		if !scanner.Scan() {
			break // Exit the loop on EOF (Ctrl+D or Ctrl+C)
		}

		// Get the input and check the domain
		checkDomain(scanner.Text())
	}

	// Check for errors in scanner
	if err := scanner.Err(); err != nil {
		log.Printf("Can't Read From Input: %v\n", err)
	}
}

// Function to check domain information
func checkDomain(domain string) {
	// Variables to store domain information
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Look up MX records for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Check if there are MX records
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Look up TXT records for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Iterate through TXT records to check for SPF
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Look up TXT records for DMARC with a specific prefix
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Iterate through DMARC records to check for DMARC configuration
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the domain information
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

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

	scanner := bufio.NewScanner(os.Stdin) // creating scanner to take input
	// from user as standard input (stdin)
	fmt.Printf("Enter an email provider (domain) : ")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // logging error if input
		//by user is not readable
		log.Fatal("Error: could not read from input:\n")
	}
}

func checkDomain(domain string) { // To check if the domain
	// provided is verifed of not

	// Declaring Variables
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord = "Sender Policy Framework not available"
	var dmarcRecord = "Domain-based Message Authentication not available"

	mxRecords, err := net.LookupMX(domain) // LookupMX returns 2 values, we are
	// assigning mxRecord and err as the 2 variables for the data

	// logging error if mxRecord is not retrievable

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(domain) // LookupTXT returns 2 values, we are
	// assigning spfRecords and err as the 2 variables for the data

	// logging error if spfrecord are not retrievable

	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	// Looping through all the spfRecords if it has prefix set to spf1,
	// spfRecords is set to spf1 if it is true

	for _, record := range spfRecords { // Iterating through records to check if
		// record matches the right prefix of spf1
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain) // LookupTXT returns 2 values,
	// we are assigning dmarcRecord and err as the 2 variables for the data

	// logging error if dmarcRecords are not retrievable

	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	// Looping through all the dmarcRecords if it has prefix set to DMARC1,
	// DMARC is set to DMARC1 if it is true

	for _, record := range dmarcRecords { // Iterating through records to check
		// if record matches the right prefix of DMARC1
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Printing all the necessary outputs

	fmt.Printf("Your Domain : %v \n", domain)
	fmt.Printf("The domain has MX (Message Exchange) : %v\n", hasMX)
	fmt.Printf("The domain has SPF (Sender Policy Framework) : %v \nHaving SPF : %v\n", hasSPF, spfRecord)
	fmt.Printf("The domain has DMARC (Domain-Based Message Authentication) : %v \nHaving DMARC : %v\n", hasDMARC, dmarcRecord)
}

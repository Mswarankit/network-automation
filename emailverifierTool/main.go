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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, spfRecord, hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	CheckErr(err)
	if len(mxRecords) > 0 {
		hasMx = true
	}
	txtRecords, err := net.LookupTXT(domain)
	CheckErr(err)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	CheckErr(err)
	for _, dmarc := range dmarcRecords {
		if strings.HasPrefix(dmarc, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = dmarc
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

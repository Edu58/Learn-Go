package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	email := "edumuriithi58@outlook.com"

	components := strings.Split(email, "@")
	_, domain := components[0], components[1]

	mx, err := net.LookupMX(domain)

	if err != nil {
		log.Fatalf("Could not Lookup MX Records for %s. Reason: %s", domain, err)
	}

	for _, record := range mx {
		fmt.Println(record)
	}
}

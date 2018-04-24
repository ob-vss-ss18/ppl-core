package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(sitename string) bool {
	fmt.Printf("Calling %s...\n", sitename)
	resp, err := http.Get(sitename)

	if err != nil {
		fmt.Printf("Failed to establish a connection.\n")
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body.\n")
		return false
	}

	fmt.Printf("Connection successful!\nResponse: %s\n\n", body)
	return true
}

func main() {
	services := [7]string{
		"https://ppl-customer.herokuapp.com/",
		"https://ppl-pricecalculator.herokuapp.com/",
		"https://ppl-billingandmailing.herokuapp.com/",
		"https://ppl-leasing.herokuapp.com/",
		"https://ppl-stock.herokuapp.com/",
		"https://ppl-reservation.herokuapp.com/",
		"https://ppl-auth.herokuapp.com/",
	}
	connected := 0
	for _, service := range services {
		if request(service) {
			connected++
		}
	}
	fmt.Print("Done.\n")
	fmt.Printf("Successful connections: %d out of 7", connected)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		in        string
		parsed_in string
	)
	fmt.Println("Do you want to remove all the data, re request it with the API and store it in database?")
	fmt.Println("write yes to confirm, any other characters to deny")
	fmt.Scan(&in)
	parsed_in = strings.ToLower(strings.Trim(in, " "))
	if parsed_in == "yes" {
		activateModel()
	} else {
		fmt.Println("Great! no unnecessary stuff!")
	}
	// next step is gocqlx
	// after that gin
	// after configurations, integrate normal CRUD operations
}

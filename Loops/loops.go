package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstname string
	var lastname string
	var email string
	var age uint

	for i := 1; i <= 2; i++ {

		fmt.Println("Enter your first name")
		fmt.Scan(&firstname)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastname)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter your age")
		fmt.Scan(&age)

		isValidName := len(firstname) > 2 && len(lastname) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidAge := age >= 15

		if isValidName && isValidEmail && isValidAge {
			fmt.Printf("Hello, %v %v\n", firstname, lastname)
			fmt.Println("Thanks for providing us your details")
		} else {
			if !isValidName {
				fmt.Println("Please enter vaild first or last name")
			}
			if !isValidEmail {
				fmt.Println("Please enter email with @")
			}
			if !isValidAge {
				fmt.Println("Please enter age greater than 15")
			}
			continue
		}
	}
}

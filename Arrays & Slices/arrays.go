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
	var bookings [5]string //Array is of fixed size and does not change themselves dynamically, so they are not flexible

	for i := 0; i < 5; i++ {
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
			bookings[i] = firstname
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
			break
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%v \n", bookings[i])
	}
}

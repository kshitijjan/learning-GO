package main

import (
	"fmt"
	"strings"
)

func getUserDetails(firstname string, lastname string, email string, age uint) (string, string, string, uint) {
	fmt.Println("Enter your first name")
	fmt.Scan(&firstname)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter your age")
	fmt.Scan(&age)

	return firstname, lastname, email, age
}
func validateDetails(firstname string, lastname string, email string, age uint) (bool, bool, bool) {
	isValidName := len(firstname) > 2 && len(lastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidAge := age >= 15

	return isValidName, isValidEmail, isValidAge
}

func main() {
	var firstname string
	var lastname string
	var email string
	var age uint
	var bookings []string //Slices are dynamic in nature, so they change themselves according to the input, Slices are recommended as compared to Arrays

	for i := 0; i < 5; i++ {

		firstname, lastname, email, age = getUserDetails(firstname, lastname, email, age)

		isValidName, isValidEmail, isValidAge := validateDetails(firstname, lastname, email, age)

		if isValidName && isValidEmail && isValidAge {
			bookings = append(bookings, firstname)
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
	for i := 0; i < 1; i++ {
		fmt.Printf("%v \n", bookings)
	}
}

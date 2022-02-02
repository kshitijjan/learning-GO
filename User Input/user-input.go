package main

import "fmt"

func main() {
	var firstname string
	var lastname string
	var email string
	var age uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstname)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter your age")
	fmt.Scan(&age)

	fmt.Printf("Hello, %v %v\n", firstname, lastname)
	fmt.Println("Thanks for providing us your details")

}

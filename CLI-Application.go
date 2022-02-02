package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var availableTickets uint = 50
var bookings []string

func greetUsers() {
	fmt.Printf("Welcome to %v ticket booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, out of which %v are available\n", conferenceTickets, availableTickets)
	fmt.Println("Get your tickets from here to attend")
}
func GetUsersDetails() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTicket uint
	fmt.Println("Enter your first name")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email adress")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)

	return firstname, lastname, email, userTicket
}

func printFirstname() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func isVaildDetails(firstname string, lastname string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstname) > 2 && len(lastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isVaildTicket := userTicket > 0 && userTicket <= availableTickets

	return isValidName, isValidEmail, isVaildTicket
}

func bookTickets(userTicket uint, firstname string, lastname string, email string) {
	availableTickets = availableTickets - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets, your tickets has been sent to %v email address\n", firstname, lastname, userTicket, email)
	bookings = append(bookings, firstname+" "+lastname)

	fmt.Printf("%v tickets are avaiable for %v\n", availableTickets, conferenceName)

}

func main() {

	greetUsers()

	for {
		firstname, lastname, email, userTicket := GetUsersDetails()
		isValidName, isValidEmail, isVaildTicket := isVaildDetails(firstname, lastname, email, userTicket)

		if isValidName && isValidEmail && isVaildTicket {

			bookTickets(userTicket, firstname, lastname, email)

			firstNames := printFirstname()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if availableTickets == 0 {
				fmt.Printf("All the tickets for %v is booked, please come next year.\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("You have entered very short first or last name")
			}
			if !isValidEmail {
				fmt.Println("You have entered an email address without '@'")
			}
			if !isVaildTicket {
				fmt.Println("You have entered invaild ticket number")
			}
			break
		}
	}

}

// Resume lecture from Packages in Go : https://youtu.be/yyUHQIec83I?t=8558

package main

import (
	"fmt"
	"strings"
)

var conferenceName = "GoLang Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	greetUser()
	for {
		firstName, lastName, email, userTickets := getUserInputs()
		input_validity := UserInputValidation(firstName, lastName, email, userTickets)
		if !input_validity {
			continue
		}
		updateRemainingTickets(firstName, lastName, userTickets)
		first_name_slice := printFirstNames()
		fmt.Printf("These are all our bookings : %v \n", first_name_slice)
		if remainingTickets <= 0 {
			fmt.Printf("All Tickets for %v have been sold\n", conferenceName)
			break
		}
	}
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets required : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func updateRemainingTickets(firstName string, lastName string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string {

	var first_name_slice []string

	for _, booking := range bookings {
		first_name_slice = append(first_name_slice, strings.Fields(booking)[0])
	}
	return first_name_slice

}

func UserInputValidation(firstName string, lastName string, email string, userTickets uint) bool {
	inputsAreValid := true
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

	if !isValidName {
		inputsAreValid = false
		fmt.Printf("Invalid First and Last Name (Each must be atleast 2 characters)\n")
	} else if !isValidEmail {
		inputsAreValid = false
		fmt.Printf("Invalid Email ID\n")
	} else if !isValidTicketCount {
		inputsAreValid = false
		fmt.Printf("Invalid Ticket Count, must be between 0 and %v\n", remainingTickets)
	}

	if inputsAreValid {
		fmt.Printf("Thankyou %v %v. You will receive an email confirmation on %v for booking %v tickets\n", firstName, lastName, email, userTickets)
	}

	return inputsAreValid

}

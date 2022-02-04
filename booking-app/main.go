// Resume lecture from arrays and slices

package main

import "fmt"

func main() {
	var conferenceName = "GoLang Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings [50]string
	bookings[0] = "Vedansh"

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for username

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address : ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets required : ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thankyou %v %v. You will receive an email confirmation on %v for booking %v tickets\n", firstName, lastName, email, userTickets)

	remainingTickets -= userTickets
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

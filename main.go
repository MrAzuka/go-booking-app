package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const totalTicketsAvailable int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	helper.WelcomeUsers(conferenceName, totalTicketsAvailable)

	for {
		var firstName string
		var lastName string
		var email string
		var numberOfTickets uint
		// Accept user input
		fmt.Println("First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Email:")
		fmt.Scan(&email)

		fmt.Println("Number of Tickets:")
		fmt.Scan(&numberOfTickets)

		// Validate user input
		isValidName, isValidEmail, validNumberOfTickets := helper.ValidateUserInput(firstName, lastName, email, numberOfTickets, remainingTickets)

		if isValidName && isValidEmail && validNumberOfTickets {
			bookTicket(firstName, lastName, email, numberOfTickets)
			wg.Add(1)
			go sendTicketToEmail(firstName, lastName, email, numberOfTickets)

			remainingTickets = remainingTickets - numberOfTickets
			if remainingTickets <= 0 {
				fmt.Println("No Available Tickets")
				// // Adding because delay needed to get the amil notification when tickets have finished
				// time.Sleep(7 * time.Second)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid Name. Cannot be less than 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email. Must contain @ symbol")
			}
			if !validNumberOfTickets {
				fmt.Println("Invalid Number Of Tickets")
			}
			fmt.Println("Invalid input. Please try again")
			continue
		}

	}
	wg.Wait()
}

func bookTicket(firstName string, lastName string, email string, numberOfTickets uint) {
	fmt.Println("Congratulations !!!!")
	fmt.Printf("%v %v your tickets have been sent to %v\n", firstName, lastName, email)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of Bookings: %v\n", bookings)
}

func sendTicketToEmail(firstName string, lastName string, email string, numberOfTickets uint) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending tickets:\n %v \n To %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}

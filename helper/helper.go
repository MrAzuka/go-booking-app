package helper

import (
	"fmt"
	"strings"
)

func WelcomeUsers(conferenceName string, totalTicketsAvailable int) {
	fmt.Printf("Welcome to the %s\n", conferenceName)
	fmt.Println("Would you like to get a ticket?")
	fmt.Printf("%v slots available\n", totalTicketsAvailable)

}

func ValidateUserInput(firstName string, lastName string, email string, numberOfTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var validNumberOfTickets bool = numberOfTickets <= remainingTickets && numberOfTickets > 0

	return isValidName, isValidEmail, validNumberOfTickets
}

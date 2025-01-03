package main

import (
	"fmt"
	"sync"
	"ticket-booking/constants"
	"ticket-booking/helper"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var bookings []UserData
var remainingTickets uint = 10
var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	userCount := 0
	for remainingTickets > 0 {

		firstName, lastName, email, city, userTickets := getUserInput()

		isValidEmail := helper.ValidateUserEmail(email)
		if !isValidEmail {
			fmt.Printf("Enter a valid email \n")
			continue
		}

		switch city {
		case "London":
			// Exce
		case "Paris":
			// Exce
		case "California":
			// Exce
		default:
			// "N"
		}

		if userTickets <= remainingTickets {
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTickets(bookings[userCount])
			userCount++
		} else {
			fmt.Printf("There are only %v tickets remaining\n", remainingTickets)
			continue
		}

		if remainingTickets == 0 {
			fmt.Printf("Ticket are sold out \n")
			break
		}
	}
	firstNamesList := printFirstNames(bookings)
	fmt.Printf("These are the tickets booked: \n %v \n", firstNamesList)
	// fmt.Printf("These are the tickets booked: \n %v \n", bookings)

	fmt.Printf("%v tickets are remaining \n", remainingTickets)
	wg.Wait()
}

/* Method to greet the users */
func greetUsers() {
	fmt.Printf("Hello, welcome to %v booking application \n", constants.ConferenceName)
	fmt.Printf("We have a total of %v tickets, and remaining of %v tickets \n", constants.TotalTickets, remainingTickets)
}

/* Method to print the first names of all the users */
func printFirstNames(bookings []UserData) []string {
	var firstNamesList []string
	for _, booking := range bookings {
		firstNamesList = append(firstNamesList, booking.firstName)
	}
	return firstNamesList
}

/* Method to get the user input */
func getUserInput() (string, string, string, string, uint) {

	var firstName, lastName, email string
	var userTickets uint
	var city string

	fmt.Printf("\n Please enter your First name and last name: ")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)

	fmt.Printf("\n Please enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("\n Enter the city you want to book your tickets in:")
	fmt.Scan(&city)

	fmt.Printf("\n Please enter the no of ticket/s: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, city, userTickets
}

/* Ticket booking logic */
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	fmt.Print("Ticket booking confirmed \n")
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
}

func sendTickets(ticketDetails UserData) {
	time.Sleep(5 * time.Second)
	fmt.Printf("\n Sending ticket details to %v", ticketDetails.firstName)
	fmt.Println("\n ####################################")
	fmt.Printf("Ticket %+v has been sent to %v", ticketDetails, ticketDetails.firstName)
	fmt.Println("\n ####################################")
	wg.Done()
}

package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	totalTickets := 50
	var userName, email string
	var remainingTickets uint = 5
	var userTickets uint
	fmt.Printf("Hello, welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and remaining of %v tickets \n", totalTickets, remainingTickets)

	fmt.Printf("Please enter your name: ")
	fmt.Scan(&userName)

	fmt.Printf("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Please enter the no of tickets: ")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		fmt.Println("Not enough tickets")
	} else {
		remainingTickets -= userTickets
		fmt.Println("Ticket booking confirmed")
		fmt.Printf("User %v has booked %v tickets. Email confirmation sent to %v \n", userName, userTickets, email)
	}

	fmt.Printf("%v tickets are remaining \n", remainingTickets)

}

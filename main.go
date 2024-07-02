package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const totalTickets = 50
	var remainingTickets = 50.6436436436436343
	fmt.Printf("Hello, welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and remaining of %v tickets \n", totalTickets, remainingTickets)
	fmt.Printf("Type of conference %T and %T\n", totalTickets, remainingTickets)

}

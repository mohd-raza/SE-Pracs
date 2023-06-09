package main

import (
	"fmt"
)
func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v conference booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickers and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend your conference")

	var bookings [50]string
	var userName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", userName, lastName, userTickets,email)
	
fmt.Printf("We had total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)}
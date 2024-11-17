package main

import "fmt" // refer to documentation for what packages contain what types of functions

func main() {
	// if you don't use a package or variable, Go will throw an error to avoid using dead code
	const conferenceName = "Go Conference"
	// for such values that do not change aka const, you can use const keyword. var keyboard allows the variable to be changed
	// if you try to change it, it will throw an error
	const conferenceTickets = 50
	var remainingTickets = 50
	// the space gets automatically added so no need to add it again
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")
	// we can also use printf to print formatted string
	// since we are not using println anymore, you have to add the \n separately for the newline
	fmt.Printf("Get your tickets below for %v\n", conferenceName)
}

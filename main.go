package main

import "fmt" // refer to documentation for what packages contain what types of functions

func main() {
	// if you don't use a package or variable, Go will throw an error to avoid using dead code
	const conferenceName string = "Go Conference"
	// for such values that do not change aka const, you can use const keyword. var keyboard allows the variable to be changed
	// if you try to change it, it will throw an error
	const conferenceTickets int = 50
	var remainingTickets int = 50
	// the space gets automatically added so no need to add it again
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")
	// we can also use printf to print formatted string
	// since we are not using println anymore, you have to add the \n separately for the newline
	fmt.Printf("Get your tickets below for %v\n", conferenceName)

	// it is important to refer to documentation about data types to figure out which data type suits the variable better depending on the usercase
	var userName string
	var userTickets int
	// initializing a variable to be filled later perhaps for user input
	// just initializing the variable will throw a warning so you need to define the data types
	// you need to tell Go Compiler the data type when declaring the variable--either by assigning a value or by declaring a type
	userName = "Tom"
	userTickets = 2
	// using placeholders
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}

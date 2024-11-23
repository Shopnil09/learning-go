package main

import "fmt" // refer to documentation for what packages contain what types of functions

func main() {
	// if you don't use a package or variable, Go will throw an error to avoid using dead code
	// syntactic sugar in programming--a term to describe a feature in a language that let's you do something more easily
	// const conferenceName string = "Go Conference"
	// instead of assigning the variable with a keyword and type, you can declare it like this. Only appropriate for 'var' and not for 'const'. Does not work if you want to explicitly want to declare a type for the variable
	conferenceName := "Go Conference"
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

	// arrays where we have to specify the size
	// var bookings [50] string
	// slices where for dynamic lists
	// what if you don't know what is the size of the array to begin with?
	// Slices--a list that is more dynamic in size. An abstraction of array
	var bookings []string
	// initializing a variable to be filled later perhaps for user input
	// just initializing the variable will throw a warning so you need to define the data types
	// you need to tell Go Compiler the data type when declaring the variable--either by assigning a value or by declaring a type

	// infinite loop
	for {
		// for fmt.Scan() we need to understand pointer. What is a pointer?
		// a pointer is a variable that points to the memory address of another variable
		// instead of passing in the userName variable which is empty, we will pass in the pointer aka memory address for the userName variable so it can store the information there
		// it is important to refer to documentation about data types to figure out which data type suits the variable better depending on the usercase
		var userName string
		var userTickets int
		var emailAddress string

		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailAddress)
		fmt.Println("Enter your ticket amount")
		fmt.Scan(&userTickets)

		// subtract the remaining ticket with the amount the user has requested to get the next remaining tickets
		remainingTickets = remainingTickets - userTickets
		// arrays and slices are common data structures in Go programming language
		// cannot mix types in this array
		bookings = append(bookings, userName)

		fmt.Printf("Thank you, %v for booking %v tickets. We will send you a confirmation to your email %v\n", userName, userTickets, emailAddress)
		fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("Whole array: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// loops are simplified in Go--you only have a for loop and you can use it for various use cases
	}
}

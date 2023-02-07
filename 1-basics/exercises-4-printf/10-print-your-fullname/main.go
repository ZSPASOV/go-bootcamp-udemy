package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  John Doe
//
// EXPECTED OUTPUT
//  Your name is John and your lastname is Doe.
// ---------------------------------------------------------

func main() {
	// v1
	firstName := os.Args[1]
	lastName := os.Args[2]
	fmt.Printf("Your name is %v and your lastname is %v.\n", firstName, lastName)
	// BONUS: Use a variable for the format specifier
	// v2
	toPrint := "Your name is %v and your lastname is %v.\n"
	fmt.Printf(toPrint, os.Args[1], os.Args[2])

	// v3
	name, lastname := os.Args[1], os.Args[2]

	msg := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, name, lastname)

}

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	numPeople := len(os.Args) - 1
	fmt.Println("There are", numPeople, "people!")
	firstName := os.Args[1]
	secondName := os.Args[2]
	thirdName := os.Args[3]
	fmt.Println("Hello great", firstName, "!")
	fmt.Println("Hello great", secondName, "!")
	fmt.Println("Hello great", thirdName, "!")
	fmt.Println("Nice to meet you all.")

	// course solution
	// var (
	// 	l  = len(os.Args) - 1
	// 	n1 = os.Args[1]
	// 	n2 = os.Args[2]
	// 	n3 = os.Args[3]
	// )

	// fmt.Println("There are", l, "people !")
	// fmt.Println("Hello great", n1, "!")
	// fmt.Println("Hello great", n2, "!")
	// fmt.Println("Hello great", n3, "!")
	// fmt.Println("Nice to meet you all.")

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Printf
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

func main() {
	// v1
	age := 33
	fmt.Printf("I'm %v years old.\n", age)
	// v2
	fmt.Printf("I'm %d years old.\n", 33)
}

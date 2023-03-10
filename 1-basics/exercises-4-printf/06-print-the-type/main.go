package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func main() {
	// v1
	fmt.Printf("Type of %v is %T\n", 3, 3)
	// v2
	fmt.Printf("Type of %d is %[1]T\n", 3)
}

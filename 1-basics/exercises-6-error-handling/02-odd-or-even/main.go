package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}
	num := os.Args[1]
	numNumeric, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("%q is not a number\n", num)
		return
	}
	if numNumeric%2 == 0 && numNumeric%8 == 0 {
		fmt.Printf("%v is an even number and it's divisible by 8\n", numNumeric)
	} else if numNumeric%2 == 0 {
		fmt.Printf("%v is an even number\n", num)
	} else {
		fmt.Printf("%v is an odd number\n", num)
	}
}

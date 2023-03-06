package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var sum int
	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || max < min {
		fmt.Println("wrong input")

		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)

		if i < max {
			fmt.Print(" + ")
		}
		sum += i
	}
	fmt.Printf(" = %v\n", sum)
}

package main

import "fmt"

func main() {
	// fallthrough statement is rarely used
	// without fallthrough statement only "big positive" will be printed and the code will stop executing
	// fallthrough executes the next clause without checking for its condition
	// fallhtrough should be the last statement in a case clause block
	i := 120
	switch {
	case i > 100:
		// fmt.Println("big positive")
		fallthrough
	case i > 0:
		fmt.Println("positive")
		fallthrough
	case i < 0:
		fmt.Println("negative")
		fallthrough
	default:
		fmt.Println("number")
	}
}

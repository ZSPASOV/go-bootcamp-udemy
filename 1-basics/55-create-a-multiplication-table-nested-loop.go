package main

import (
	"fmt"
	"os"
	"strconv"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

// const max = 5

func main() {
	max, _ := strconv.Atoi(os.Args[1])

	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// multiplication
	fmt.Println("multiplication:")
	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
	fmt.Println("addition:")
	// addition
	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i+j)
		}
		fmt.Println()
	}
	fmt.Println("sbtrkt:")
	// addition
	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}

}

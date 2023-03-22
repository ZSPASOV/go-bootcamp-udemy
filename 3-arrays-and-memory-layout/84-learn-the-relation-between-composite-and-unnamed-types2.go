package main

import "fmt"

func main() {
	type bookcase [5]int
	// declaring two arrays using unnamed types
	blue := [5]int{6, 9, 3, 2, 1}
	red := [5]int{6, 9, 3, 2, 1}
	//declaring one array with named type bookcase
	green := bookcase{6, 9, 3, 2, 1}

	fmt.Println("Are they equal?")
	if blue == red {
		fmt.Println("✓")
	} else {
		fmt.Println("✖")
	}

	if blue == green {
		fmt.Println("✓")
	} else {
		fmt.Println("✖")
	}

	// noticed that only the green array has a named type
	fmt.Printf("blue: %#v\n", blue)   // blue: [5]int{6, 9, 3, 2, 1}
	fmt.Printf("red: %#v\n", red)     // red: [5]int{6, 9, 3, 2, 1}
	fmt.Printf("green: %#v\n", green) // green: main.bookcase{6, 9, 3, 2, 1}
}

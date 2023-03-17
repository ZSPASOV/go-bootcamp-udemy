package main

import "fmt"

func main() {
	// arrays are copied
	blue := [3]int{6, 9, 3}
	red := blue
	// they are equal, but are in different locations in memory
	fmt.Println(blue == red)
	blue[0] = 10
	fmt.Println(blue) // [10 9 3]
	fmt.Println(red)  // [6 9 3]

	// changing one of them wont affect the other
}

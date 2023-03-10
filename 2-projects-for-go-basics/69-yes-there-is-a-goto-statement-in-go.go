package main

import "fmt"

// goto is rarely used so this lecture is optional

func main() {
	var i int

	// this works the same as loop
loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}
	fmt.Println("done")
}

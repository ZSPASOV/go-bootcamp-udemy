package main

import (
	"fmt"
	"math/rand"
)

func main() {

	guess := 10

	// that is a pseudo random generator, it will generate always the same numbers
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

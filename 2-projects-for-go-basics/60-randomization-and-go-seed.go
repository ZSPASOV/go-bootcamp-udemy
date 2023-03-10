package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rerun with different value of seed and different numbers will be generated
	rand.Seed(28)

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

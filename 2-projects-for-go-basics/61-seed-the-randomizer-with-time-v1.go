package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	timeSeed := t.UnixNano()
	rand.Seed(timeSeed)

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

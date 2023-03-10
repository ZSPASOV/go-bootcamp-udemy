package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
	search:
		for i, w := range words {

			switch q {
			case "and", "or", "the":
				break search // you can use a labeled break to quit from the loop, instead of the switch
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

			}
		}
	}
}

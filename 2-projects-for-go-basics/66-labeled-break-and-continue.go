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

queries: // the label of the parent loop
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break queries // labeled break statement in that case it is needed to break the parent loop. if there was no label, the nested loop would break.
			}
		}
	}
}

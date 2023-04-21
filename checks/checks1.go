package main

import (
	"fmt"
	"strings"
)

const words = "ny, la, rio, sao paulo"

func main() {
	a := []string{"pld", "sf", "varna"}
	for i, v := range a {
		fmt.Println(i)
		fmt.Println(v)
	}

	wordsSlice := strings.Fields(words)

	for i, v := range wordsSlice {
		fmt.Println(i)
		fmt.Println(v)
	}
}

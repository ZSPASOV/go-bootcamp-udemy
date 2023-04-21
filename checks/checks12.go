package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const word = "gökyüzü"
	bword := []byte(word)

	// ö => 2 bytes
	// ü => 2 bytes
	fmt.Println(utf8.RuneCount(bword), len(word), len(string(word[1])))
	for _, val := range bword {
		fmt.Println(val)
	}
	fmt.Println(word[1])
}

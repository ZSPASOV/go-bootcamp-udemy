package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "asd"
	wordB := []byte(word)
	fmt.Println(wordB)
	rune, size := utf8.DecodeRune(wordB)
	fmt.Println(rune)
	fmt.Println(size)
	fmt.Println(string(wordB))
}

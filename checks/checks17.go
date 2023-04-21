package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const nihongo = "日本語"
	fmt.Println(len(nihongo))
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	fmt.Println("With DecodeRuneInString:")
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

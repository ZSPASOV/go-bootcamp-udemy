package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const text = `Според мен Юнайтед трябва да се борят за върха, както беше в миналото. Мисля, че им трябва централен нападател. Някой, който да носи голове.
Купуваш си централен нападател и казваш, че това е човекът, който ще носи по 25 гола. Нуждаят се от такъв играч. Трябва да потърсят класен голмайстор.
Головата статистика на Хари Кейн е безспорна. `

	r, size := utf8.DecodeRuneInString("Köln")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("öln")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ln")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("n")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	// for range loop automatically decodes the runes
	//   but it gives you the position of the current rune
	//   instead of its size.

	// for _, r := range text {}
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		i += size
	}
	fmt.Println()
}

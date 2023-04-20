package main

import "fmt"

func main() {
	str := "Yūgen ☯ 💀"

	// can't change a string
	// a string is a read-only byte-slice
	// str[0] = 'N'
	// str[1] = 'o'

	// that creates a []byte and copies the bytes of the string to the new slice's backing array
	bytes := []byte(str)
	fmt.Printf("%#v\n", bytes)
	fmt.Println(bytes)
	// for _, val := range bytes {
	// 	fmt.Println(val)
	// 	fmt.Printf("%c\n", val)
	// }

	// you can change a byte slice
	bytes[0] = 'N'
	bytes[1] = 'o'
	fmt.Printf("%s\n", bytes)
}

package main

import "fmt"

func main() {
	// in the unicode standard there is a million of code points, for now we will print just a tiny part of them
	//rune literals are typeless. They can be assigned to a variable with any numeric type. each gets converted to int here
	start, stop := 'A', 'Z'

	fmt.Println(start, stop)
	// decimal
	for n := start; n <= stop; n++ {
		// %c verb prints a character depending on the given code point
		fmt.Printf("%c => %[1]d\n", n)
	}

	// hexadecimal
	for n := start; n <= stop; n++ {
		// %c verb prints a character depending on the given code point
		fmt.Printf("%c => %[1]x\n", n)
	}

}

package main

import "fmt"

func main() {

	// Array literal
	// creates & initializes a new array with the given values
	var books = [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	fmt.Printf("%#v\n", books)

	// short declaration example

	booksTwo := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	// no need of comma when the elements are on the same line
	booksThree := [2]string{"Vonegut", "Orwel"}
	fmt.Printf("%#v\n", booksTwo)
	fmt.Printf("%#v\n", booksThree)
}

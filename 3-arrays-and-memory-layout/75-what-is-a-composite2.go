package main

import "fmt"

func main() {
	booksOne := [4]string{"one",
		"two",
		"three",
		"four"}
	fmt.Printf("%#v\n", booksOne) // [4]string{"one", "two", "three", "four"}

	booksTwo := [4]string{"one",
		"two",
	}
	fmt.Printf("%#v\n", booksTwo) // [4]string{"one", "two", "", ""}

	// go finds out the length of the array automatically - the ELLIPSIS operator
	booksThree := [...]string{"one",
		"two",
	}
	fmt.Printf("%#v\n", booksThree) // [2]string{"one", "two"}
}

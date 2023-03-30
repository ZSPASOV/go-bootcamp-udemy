package main

import "fmt"

func main() {
	var books [5]string
	books[1] = "Dracula"
	books[3] = "1984"

	var games []string

	fmt.Printf("books : %#v\n", books)
	fmt.Printf("the length  of games : %d\n", len(games))

	// a slice can only be compared to a nil value
	fmt.Printf("games : %t\n", games == nil) // true

	otherGames := []string{"pacman", "doom"}
	fmt.Println(otherGames == nil) // false

	// fmt.Println(games == otherGames) // invalid operation: cannot compare games == otherGames (slice can only be compared to nil)

	// one of the ways to compare slices is with a loop
	var ok string

	moreGames := []string{"mgs", "dmc"}

	for i, game := range otherGames {
		if game != moreGames[i] {
			ok = "not"
			break
		}
	}
	fmt.Printf("otherGames and moreGames are %s equal\n", ok)

	// slices with the same element type can be assigned to each other, no matter what their lengths are
	games = moreGames
	fmt.Printf("%#v\n", games) //[]string{"mgs", "dmc"}
	// we can assign a nil value to a slice
	games = nil
	fmt.Printf("%#v\n", games) //[]string(nil)
	games = []string{}
	fmt.Printf("%#v\n", games) //[]string{}

	// an empty slice is not a nil slice
	// an empty slice is an initialized slice
	// a nil slice is an uninitialized slice
	// their lengths are the same: Zero
}

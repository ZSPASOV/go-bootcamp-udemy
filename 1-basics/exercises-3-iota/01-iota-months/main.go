package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------

func main() {
	// const (
	// 	Nov = 11
	// 	Oct = 10
	// 	Sep = 9
	// )
	// v1
	// const (
	// 	Nov = -(- 11 + iota)
	// 	Oct
	// 	Sep
	// )
	// v2
	const (
		Nov = 11 - iota // 11 - 0 = 11
		Oct             // 11 - 1 = 10
		Sep             // 11 - 2 = 9
	)

	fmt.Println(Sep, Oct, Nov)
}

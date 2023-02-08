package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool
	// default value is false
	fmt.Printf("%v\n", big)

	// if radius >= 50 {
	// 	if radius >= 100 {
	// 		if radius >= 200 {
	// 			big = true
	// 		}
	// 	}
	// }

	// if big != true {
	// 	fmt.Println("I don't know.")
	// } else if !(isSphere == false) {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }

	// my solution
	if radius >= 200 {
		big = true
	}

	if big && isSphere {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}

	// course solution :

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

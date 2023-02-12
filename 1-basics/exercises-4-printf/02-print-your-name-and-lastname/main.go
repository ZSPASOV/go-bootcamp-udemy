package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is firstName and my lastname is lastName.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier

	// fmt.Printf("?", ?, ?)
	firstName := "Zapryan"
	lastName := "Spasov"

	fmt.Printf("My name is %v and my lastname is %v.\n", firstName, lastName)
	fmt.Printf("My name is %s and my lastname is %s.\n", "Zapryan", "Spasov")

	// BONUS
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Zapryan", "Spasov")
}

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	dict1 := map[string]string{
		"Spasov":   "53035209",
		"Stoev":    "21421",
		"Georgiev": "35515135",
	}
	fmt.Printf("%#v\n", dict1)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	dict2 := map[int]bool{
		2: true,
		5: false,
		4: true,
		7: true,
	}
	fmt.Printf("%#v\n", dict2)
	// #3
	// Key        : Last name
	// Element    : Phone numbers
	dict3 := map[string][]string{
		"Spasov": []string{
			"35235253532",
			"3525523"},
		"Stoev": []string{
			"532532",
			"61616"},
		"Georgiev": []string{
			"421421",
			"531513"},
	}
	fmt.Printf("%#v\n", dict3)
	// the same as dict3, but with the simplified syntax (the correct one)
	dict4 := map[string][]string{
		"Spasov": {
			"35235253532",
			"3525523"},
		"Stoev": {
			"532532",
			"61616"},
		"Georgiev": {
			"421421",
			"531513"},
	}
	fmt.Printf("%#v\n", dict4)
	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	dict5 := map[string]map[string]int{
		"Spasov": {
			"Bananas": 5,
			"Oranges": 8,
			"Apples":  2},
		"Stoev": {
			"Bananas": 5,
			"Pears":   7,
			"Oranges": 8,
			"Apples":  2},
	}

	fmt.Printf("%#v\n", dict5)
}

package main

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	var (
		names = [...]string{
			"Kiro",
			"Ivan",
			"Niki",
		}
		distances = [...]int{
			10,
			25,
			40,
			67,
			67,
		}
		data = [...]byte{
			'A',
			'B',
			2,
			4,
			'C',
		}
		ratios = [...]float64{
			4.7,
		}
		alives = [...]bool{
			true,
			true,
			false,
			false,
		}
		zero [0]byte // A byte array that doesn't occupy memory space
	)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range distances {
		fmt.Println(value)
	}

	for _, value := range data {
		fmt.Println(value)
	}
	for _, value := range ratios {
		fmt.Println(value)
	}
	for _, value := range alives {
		fmt.Println(value)
	}
	fmt.Println(zero)
}

package main

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
		names = [3]string{
			"Kiro",
			"Ivan",
			"Niki",
		}
		distances = [5]int{
			10,
			25,
			40,
			67,
			67,
		}
		data = [5]byte{
			'A',
			'B',
			2,
			4,
			'C',
		}
		ratios = [1]float64{
			4.7,
		}
		alives = [4]bool{
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

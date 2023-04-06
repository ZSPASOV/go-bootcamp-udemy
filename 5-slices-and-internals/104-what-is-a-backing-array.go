package main

import "fmt"

func main() {
	// a slice is a window to its backing array
	// slice literal creates a hidden array and returns a slice that refers to that array
	ages := []int{35, 15, 25}
	// a slice expression does not create a new backing array, it just refers to some part of the original backing array
	fmt.Println(ages[1:3])

	// slicing is a fast and cheap operation, because of the same backing array is used

	// a sliced array becomes the backing array of the returned slice
	agesArray := [3]int{35, 15, 25}
	newAges := agesArray[0:3]
	fmt.Println(newAges)
	newAges[0] = 100
	newAges[1] = 50
	fmt.Println(agesArray)
	fmt.Println(newAges)
}

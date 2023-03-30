package main

import "fmt"

func main() {
	// its length is part of the type of the array
	// the zero value of an array is the zero valued elements
	// an array can contain only the same type of elements
	var numsArray [5]int
	fmt.Println(len(numsArray))    // 5
	fmt.Printf("%#v\n", numsArray) // [5]int{0, 0, 0, 0, 0}

	// its length is not part of the type of the slice. it can grow and shrink
	// the zero value of a slice is nil
	// a slice can contain only the same type of elements
	var numsSlice []int
	fmt.Println(len(numsSlice))    // 0
	fmt.Printf("%#v\n", numsSlice) // []int(nil)

}

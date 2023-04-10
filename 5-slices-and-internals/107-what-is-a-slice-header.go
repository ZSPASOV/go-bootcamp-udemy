package main

import "fmt"

func main() {
	// a slice header is a data structure.
	// a slice header has:
	// a pointer - it stores the memory location of the slice's backing array. it is an integer where the RAM location of the first element of the backing array is
	// length - it looks for the exact number of elements of the backing array
	// capacity - the number of fields the backing array actualy has, starting from the first selected element
	agesArray := [9]int{
		10,
		15,
		20,
		35,
		55,
		200,
		100,
		500,
		700,
	}

	// the cap function returns the capacity
	agesSlice := agesArray[2:4]
	fmt.Println(cap(agesSlice)) // 7 - it returns 7, because it starts from the second the third location of the backing array. the previous 2 elements become invisible.
	agesAnotherSlice := agesArray[0:2]
	fmt.Println(cap(agesAnotherSlice)) // 9 - the length of agesSlice and agesAnotherSlice is the same, but the capacity is 9, because it starts from the 0 position of the backing array

	// a nil slice does not have a backing array, but it still has a slice header
	var agesNil []int
	fmt.Println(cap(agesNil)) // 0
}

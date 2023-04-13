package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// make allows you to pre-allocate a backing array of a slice with a given length and capacity
	// it initializes and returns a slice with the given legnth and capacity
	// we use the make function because using a slice literal, if a slice needs to grow and if there is no enough capacity, the append function allocates a new backing array.
	// that is an expensive operation. you can prevent that by prealocating a large enough backing array using the make function

	// if you do not pass a third argument, the length and the capacity is the same
	sliceA := make([]int, 2)
	s.Show("SliceA:", sliceA) // SliceA:             (len:2  cap:2  ptr:1456)
	// you can initialize a slice with a different length and capacity
	sliceB := make([]int, 3, 5)
	s.Show("sliceB: ", sliceB) // sliceB:             (len:3  cap:5  ptr:4304)

	sliceC := make([]int, 0, 5)
	s.Show("sliceC before append: ", sliceC)
	sliceC = append(sliceC, 42)
	s.Show("sliceC after append: ", sliceC)
}

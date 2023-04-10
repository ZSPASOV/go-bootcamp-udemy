package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	// when the capacity is full, append() allocates a new and larger array
	ages := []int{35, 15}
	s.Show("ages: ", ages)
	ages = append(ages, 5)
	s.Show("ages: ", ages)
}

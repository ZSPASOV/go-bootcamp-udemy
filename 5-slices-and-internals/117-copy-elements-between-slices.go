package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// copy - copies the elements of a slice to another slice
	evens := []int{2, 4}
	odds := []int{3, 5, 7}
	copy(evens, odds)
	// the copy copies based on length
	s.Show("evens:", evens)

	data := []float64{10, 25, 30, 50}
	s.Show("Probabilities", data)
}

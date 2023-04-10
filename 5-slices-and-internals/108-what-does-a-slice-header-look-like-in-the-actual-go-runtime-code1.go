package main

import (
	s "github.com/inancgumus/prettyslice"
)

type collection [4]string

func main() {
	// this is the implementation of the slice data type in the actual go runtime source-code
	// as you can see it is a data structure with 3 integer fields:

	// https://go.dev/src/runtime/slice.go
	// type slice struct {
	// 	array unsafe.Pointer
	// 	len   int
	// 	cap   int
	// }

	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	// the original values of the data array do not change.
	// only the change data value changes. It is because passing a value to a function makes a new copy of it.
	change(data)
	s.Show("main data", data)
}

func change(data collection) {
	data[2] = "brilliant!"
	s.Show("changed data", data)
}

package main

import (
	s "github.com/inancgumus/prettyslice"
)

type collection []string

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
	// same as the previous version, however type collection is a slice now.
	// now both data (in main and the passed in change function) change their values.
	// this is because slice value = slice header.
	// Go makes a new copy of the slice header and passes it to
	// the function. But it does not make a copy of the backing array.
	// So, the slices still use the same backing array.
	change(data)
	s.Show("main data", data)
}

func change(data collection) {
	data[2] = "brilliant!"
	s.Show("changed data", data)
}

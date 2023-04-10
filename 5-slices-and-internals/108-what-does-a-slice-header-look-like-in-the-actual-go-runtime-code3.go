package main

import (
	"fmt"
	"unsafe"

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

	array := [...]string{"slices", "are", "awesome", "period"}

	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(array)) // array's size: 64 bytes
	fmt.Printf("slice's size: %d bytes\n", unsafe.Sizeof(data))  // slice's size: 24 bytes
	// even though the array and the data slice have the same element, their size is different.
	// the size of the slice is smaller, because it only stores 3 integer fields in a slice header. so on a 64 bit machine, this means 24 bytes.
	// so the size of a slice is always fixed to 24 bytes. however, an array's size changes depending on its elements

	// SO IN SUMMARY - creating, assigning, slicing and passing a slice to a function is very cheap operation, because it is a very cheap operation
	// because of the fixed size of the slice. however it is not the same with an array. it is because the array will be coppied along with
	// all of its items. when you pass a slice, only its header will be copied. also the slice does not carry with it its backing array.
	// it only refers to it with its pointer field in the memory.
}

func change(data collection) {
	data[2] = "brilliant!"
	s.Show("changed data", data)
}

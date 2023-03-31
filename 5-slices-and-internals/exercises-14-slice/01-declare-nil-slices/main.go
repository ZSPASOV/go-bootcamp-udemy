package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	var (
		names     []string
		distances []float64
		data      []uint8
		ratios    []float64
		alives    []bool
	)

	fmt.Printf("Type of names: %T, length of names : %v, equal to nil %t\n", names, len(names), names == nil)
	fmt.Printf("Type of distances: %T, length of distances : %v, equal to nil %t\n", distances, len(distances), distances == nil)
	fmt.Printf("Type of data: %T, length of data : %v, equal to nil %t\n", data, len(data), data == nil)
	fmt.Printf("Type of ratios: %T, length of ratios : %v, equal to nil %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("Type of alives: %T, length of alives : %v, equal to nil %t\n", alives, len(alives), alives == nil)
}

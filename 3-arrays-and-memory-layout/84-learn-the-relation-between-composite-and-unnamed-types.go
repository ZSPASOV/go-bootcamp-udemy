package main

import "fmt"

func main() {
	// named type
	type bookcase [3]int

	// compare it with the unnamed type on the right
	fmt.Println(bookcase{6, 9, 3} == [3]int{6, 9, 3}) // true

	// they are equal, because their underlying types are identical
	// in that case the underlying type is [3]int

	// compare named with named
	type cabinet [3]int

	// fmt.Println(bookcase{6, 9, 3} == cabinet{6, 9, 3})
	// 3-arrays-and-memory-layout/84-learn-the-relation-between-composite-and-unnamed-types.go:18:35: invalid operation: bookcase{…} == cabinet{…} (mismatched types bookcase and cabinet)
	// we cannot compare them because their types are different. a named type is always a different type than any other type.

	// the only way to compare them is if we convert one of them to the type of the other
	fmt.Println(bookcase{6, 9, 3} == bookcase(cabinet{6, 9, 3})) // true
	fmt.Println(cabinet(bookcase{6, 9, 3}) == cabinet{6, 9, 3})  // true
}

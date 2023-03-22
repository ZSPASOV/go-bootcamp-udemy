package main

import "fmt"

func main() {
	// v1
	nestedArr := [2][3]int{
		[3]int{5, 6, 1},
		[3]int{9, 8, 4},
	}
	fmt.Println(nestedArr)
	// v2 - go allows you to omit the types of the inner arrays
	nestedArr2 := [2][3]int{
		{5, 6, 1},
		{9, 8, 4},
	}
	fmt.Println(nestedArr2)
}

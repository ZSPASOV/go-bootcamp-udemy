package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// remember: case condition's types should match to condition expresion's type
	// in go the code is read from top to bottom, so no need to write anything for default statement
	// because the only option for it is to be zero
	// v1
	i, _ := strconv.Atoi(os.Args[1])
	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
	// same as: in go switch condition by default is true, so above we did not need to write switch true
	// but to see the full sytntax see below :
	// v2
	switch true {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
	// v3 - if default is at the beginning
	switch {
	default:
		fmt.Println("zero")
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	}
}

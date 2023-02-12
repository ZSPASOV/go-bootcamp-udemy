package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number : ", n)
	fmt.Println("Converted number : ", err)
}

// func Atoi(s string) (int, error) {
//  // returns a value with an error type. success -> nil (uninitialized error value)
//  // error -> non-nill (initialized error value)
// }

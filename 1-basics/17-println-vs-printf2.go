package main

import "fmt"

func main() {
	operations := 2350
	success := 543
	fail := 433
	name := "Zapryan"
	distance := 25.5
	yeah := true
	fmt.Printf("total: %d success: %d / %d\n", operations, success, fail)
	// printing types
	fmt.Printf("%T\n", operations)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", distance)
	fmt.Printf("%T\n", distance)
	fmt.Printf("%x\n", distance)
	fmt.Printf("%X\n", distance)
	fmt.Printf("%v\n", distance)
	fmt.Printf("%#v\n", distance)
	fmt.Printf("%t\n", yeah)
	// println escape sequences
	fmt.Println("hi")
	fmt.Println("hi\\n\"hi\"")
}

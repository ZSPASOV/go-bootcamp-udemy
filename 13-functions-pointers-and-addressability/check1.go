package main

import "fmt"

func incr(a int) int {
	a++
	return a
}

func main() {
	num := 10
	fmt.Println(incr(num))
}

package main

import "fmt"

func main() {
	blue := [3]int{6, 9, 3}
	red := [3]int{6, 9, 3}
	yellow := [5]int{6, 9, 3}
	green := [3]int{9, 6, 3}
	fmt.Println(blue == red)
	fmt.Println(green == blue)
	fmt.Printf("yellow array: %#v\n", yellow)
}

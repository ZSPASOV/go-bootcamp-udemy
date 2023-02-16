package main

import "fmt"

func main() {
	var (
		sum int
		i   = 1
	)
	// this is an infinite loop:
	// for {
	// 	sum += i
	// 	i++
	// 	fmt.Println(sum)
	// }

	// you can exit from the loop by using the break statement
	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
}

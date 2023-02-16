package main

import "fmt"

func main() {
	var (
		sum int
		i   = 1
	)
	// you can skip the rest of the current step by using the continue statement
	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
}

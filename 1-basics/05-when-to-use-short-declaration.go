package main

import "fmt"

// normal 2:
var version string

func main() {
	// normal 1:
	var score int // already score = 0
	// normal 3:
	var (
		video    string
		duration int
		current  int
	)
	// short 1, 2:
	width, height := 100, 50

	// short 3:
	// redeclaration is both blessing and a curse. below width is redeclared
	width, color := 50, "red"
	fmt.Println(score, video, duration, current, width, height, color)
}

// use normal declaration when:
// 1. You do not know the initial value
// 2. When you need a package scoped variable
// 3. When you want to group variables for greater readability

// Short is the most used and prefered declaration in GO. Use short declaration when:
// 1. If you know the initial value
// 2. Keep the code short and concise
// 3. For redeclaration

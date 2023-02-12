package main

import "fmt"

func main() {
	var score int
	score = 4
	var valid bool
	valid = true
	var notValid bool
	notValid = false

	if score > 3 && valid == true {
		fmt.Println("good")
	}
	// v2 you can use bool values directly:
	if score > 3 && valid {
		fmt.Println("good")
	}
	if score > 3 && !notValid {
		fmt.Println("not good")
	}
}

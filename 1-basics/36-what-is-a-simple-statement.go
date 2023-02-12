package main

import (
	"fmt"
	"strconv"
)

// simple statements are completely optional
// simple statements allow you to execute a statement inside another statement
// the semicolon ; separator separates the simple statement and the condition expression
// declared variables are only visible inside the if statement (and its branches)
func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Printf("There was no error, n is", n)
	}
}

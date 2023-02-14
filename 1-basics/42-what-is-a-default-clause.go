package main

import (
	"fmt"
)

func main() {
	// similar to an if statement's else branch
	// the default clause is executed when no cases match
	city := "Limbo"
	switch city {
	case "Paris":
		fmt.Println("France")
	default:
		fmt.Println("Where")
	}
}

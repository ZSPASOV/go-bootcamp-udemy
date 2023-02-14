package main

import (
	"fmt"
	"os"
)

func main() {
	// similar to logigal or ||
	// you can add as many conditions as you want (keep it readable)
	city := os.Args[1]
	switch city {
	case "Paris", "Lyon", "Marseille":
		fmt.Println("France")
	case "Tokyo", "Osaka":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}
	// it is the same as :
	if city == "Paris" || city == "Lyon" || city == "Marseille" {
		fmt.Println("France")
	} else if city == "Tokyo" || city == "Osaka" {
		fmt.Println("Japan")
	}
}

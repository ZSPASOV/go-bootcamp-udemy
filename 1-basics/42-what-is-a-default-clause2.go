package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]
	// you can move the default clause anywhere inside the switch statement
    // there can be only one default clause per switch statement
	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}
}

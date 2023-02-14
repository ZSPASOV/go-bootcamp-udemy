package main

import (
	"fmt"
	"os"
)

func main() {
	// the type of the case conditions depends on the type of the city variable. if it is string, the case conditions should be string also.
	// values should be unique among case conditions.
	// each case clause creates its own unique block
	city := os.Args[1]
	switch city {
	case "Paris":
		fmt.Println("France")
		// break - the break statements are optional, Go does that behind the scenes, so it is not mandatory to use break
		// the vip variable is only available inside this case block :
		vip := true
		fmt.Println("Vip trip?", vip)
	case "Tokyo":
		fmt.Println("Japan")
		// break - the break statements are optional, Go does that behind the scenes, so it is not mandatory to use break
	}
	// it is the same as :
	if city == "Paris" {
		fmt.Println("France")
	} else if city == "Tokyo" {
		fmt.Println("Japan")
	}
}

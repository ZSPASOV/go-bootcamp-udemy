package main

import (
	"fmt"
	"time"
)

func main() {
	// time.now() returns the current time as Time
	// .Hour( gets the hour as int from the Time instance)
	t := time.Now()
	fmt.Printf("%T\n", t)
	hour := t.Hour()
	switch {
	case hour >= 18:
		fmt.Println("good evening")
	case hour >= 12:
		fmt.Println("good afternoon")
	case hour >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")
	}
}

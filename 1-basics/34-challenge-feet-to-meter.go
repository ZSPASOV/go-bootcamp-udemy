package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	feet := os.Args[1]
	feet_int, err := strconv.Atoi(feet)
	if err != nil {
		fmt.Printf("error: %v is not a number\n", feet)
		return
	}
	meters := float64(feet_int) * 0.3048
	fmt.Printf("%v feet is %.2f meters.", feet, meters)
}

package main

import "fmt"

func main() {
	msg := []string{"h", "e", "l", "l", "o"}
	fmt.Println(msg[1:3])
	fmt.Println(msg[:5])
	fmt.Println(msg[:4])
	fmt.Println(msg[:])
	fmt.Println(msg[2:4])
	fmt.Println(msg[1:])

	newSlice := append(msg[1:4], "!")
	fmt.Println(newSlice)
	fmt.Println(msg)
}

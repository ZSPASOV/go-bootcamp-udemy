package main

import "fmt"

func main() {
	randomString := "keeper"
	for _, char := range randomString {
		// fmt.Printf("%c\n", char)
		fmt.Println(char)
	}
}

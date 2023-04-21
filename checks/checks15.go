package main

import "fmt"

func main() {
	word := "Panyo"
	fmt.Println(word)
	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
	}
	fmt.Println(word[2])
}

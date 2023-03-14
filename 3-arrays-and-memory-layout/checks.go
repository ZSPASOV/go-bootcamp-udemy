package main

import "fmt"

func main() {
	var names [3]string

	names[len(names)-1] = "!"
	names[1] = "think" + names[2]
	names[0] = "Don't"
	names[0] += " "

	fmt.Println(names[0] + names[1] + names[2])
}

package main

import "fmt"

func main() {
	var todo []string
	todo = append(todo, "sing")
	todo = append(todo, "run")
	todo = append(todo, "code", "play")

	tomorrow := []string{"see mom", "learn go"}

	todo = append(todo, tomorrow...)

	fmt.Println(todo)
}

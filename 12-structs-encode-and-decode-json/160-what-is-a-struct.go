package main

import "fmt"

func main() {
	type VideoGame struct {
		Title     string
		Genre     string
		Published bool
	}

	pacman := VideoGame{
		Title:     "Pac-Man",
		Genre:     "Arcade Game",
		Published: true,
	}

	fmt.Println(pacman)
}

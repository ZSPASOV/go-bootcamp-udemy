package main

import "fmt"

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	top3 := items[:3]
	last4 := items[9:]

	fmt.Println(top3)
	fmt.Println(last4)
}

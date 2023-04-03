package main

import "fmt"

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	const pageSize = 4

	l := len(items)

	for i := 0; i < l; i += pageSize {
		to := i + pageSize
		if to > l {
			to = l
		}

		currentPage := items[i:to]

		// .Sprintf returns a formatted string
		head := fmt.Sprintf("Page %d :", (i/pageSize)+1)
		fmt.Println(head)
		fmt.Println(currentPage)
	}
}

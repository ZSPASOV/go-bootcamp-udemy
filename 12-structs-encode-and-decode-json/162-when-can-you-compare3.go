package main

import "fmt"

type song struct {
	title, artist string
}

func main() {
	finalHome := song{
		title:  "Final Home",
		artist: "DJ Krush",
	}

	innerCityLife := song{
		title:  "Inner City Life",
		artist: "Goldie",
	}

	songs := []song{
		finalHome,
		innerCityLife,
	}

	for _, song := range songs {
		fmt.Printf("%#v\n", song)
	}
}

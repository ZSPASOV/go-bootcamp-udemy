package main

import "fmt"

func main() {
	type song struct {
		Title  string
		Artist string
	}

	firstSong := song{
		Title:  "Wonderwall",
		Artist: "Oasis",
	}
	secondSong := song{
		Title:  "Wonderwall",
		Artist: "Oasis",
	}

	thirdSong := song{
		Title:  "Other song",
		Artist: "Oasis",
	}

	fmt.Println(firstSong == secondSong)
	fmt.Println(firstSong == thirdSong)
}

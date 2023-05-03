package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	// anonymous field text here, the type is the type of text
	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{
			title: "moby dick",
			words: 206052,
		},
		isbn: "102030",
	}

	// v1
	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.text.title, moby.text.words, moby.isbn)
	// v2
	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, moby.words, moby.isbn)
}

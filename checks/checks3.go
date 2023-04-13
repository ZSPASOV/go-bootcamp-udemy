package main

import s "github.com/inancgumus/prettyslice"

func main() {
	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	s.Show("WOrds", words)
	words = append(words[:5], "crystals", "and", "diamonds")
	s.Show("WOrds", words)
}

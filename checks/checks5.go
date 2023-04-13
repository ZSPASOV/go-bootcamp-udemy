package main

import s "github.com/inancgumus/prettyslice"

func main() {
	words := []string{1022: ""}
	words = append(words, "boom!")
	s.Show("words: ", words)
}

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	words := []string{1023: ""}
	words = append(words, "boom!")
	s.Show("Words: ", words)
}

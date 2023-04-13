package main

import s "github.com/inancgumus/prettyslice"

func main() {
	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = append(words[:1], "is", "everywhere")
	words = append(words, words[len(words)+1:cap(words)]...)
	s.Show("werds: ", words)
}

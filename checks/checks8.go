package main

import s "github.com/inancgumus/prettyslice"

func main() {
	lyric := []string{"show", "me", "my", "silver", "lining"}
	s.Show("lyric: ", lyric)
	part := lyric[:2:2]
	s.Show("part before append", part)
	part = append(part, "right", "place")
	s.Show("part after append", part)
}

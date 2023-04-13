package main

import s "github.com/inancgumus/prettyslice"

func main() {
	lyric := []string{"show", "me", "my", "silver", "lining"}
	s.Show("lyric", lyric)
	part := lyric[1:3:5]
	s.Show("part", part)
}

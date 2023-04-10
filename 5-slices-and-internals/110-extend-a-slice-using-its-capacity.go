package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// shows the backing array of the slice:
	s.PrintBacking = true

	// below is a nil string slice
	// it does not have a backing array, but it has a slice header
	var games []string

	// when a slice value does not have a backing array, its pointer field will be zero => nil => uninitialized
	// when a slice value does not have a pointer, its capacity field will be zero.
	// a slice's capacity can be zero, but it can have a pointer
	s.Show("the games: ", games) // the games:          (len:0  cap:0  ptr:0   )

	// let's overwrite the nil slice from above, with an initialized empty slice
	// this time its pointer is not 0, the capacity is 0
	// GO assigns the same dummy backing array to all the empty slices
	games = []string{}
	s.Show("the unitialized games: ", games) // the unitialized games:   (len:0  cap:0  ptr:9304)

	games = []string{"pacman", "mario", "tetris", "doom"}
	// note - for the non-empty games slice, the ptr: value will be different each time we rerun the script
	s.Show("non-empty games: ", games) // non-empty games:    (len:4  cap:4  ptr:1792)

	// the pointers of these slices are the same. they both are looking at the same backing array
	part := games
	s.Show("the part:", part) // non-empty games:    (len:4  cap:4  ptr:1792)

	part = games[:0]
	s.Show("part[:0]", part) // part[:0]            (len:0  cap:4  ptr:1792)

	// we can reslice to capacity:
	s.Show("part[:cap]", part[:cap(part)]) // (len:4  cap:4  ptr:1792)

	// an example with loop:
	// note that, because a string value is exactly 16 bytes on a 64 bit machine, the pointer value incereases with 16

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

}

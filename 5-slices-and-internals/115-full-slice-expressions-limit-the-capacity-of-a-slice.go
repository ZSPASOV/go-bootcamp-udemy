package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// a full slice expression allows you to control the capacity of the returned slice from a slice expression
	// sometimes you need to limit yourself
	// newSlice := slicable[START:STOP:CAP]
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := nums[:2:2]
	s.Show("nums", nums)
	s.Show("odds before append", odds)
	odds = append(odds, 5, 7)
	s.Show("odds after append", odds)
}

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	var nums []int
	s.Show("no backing array:", nums)

	nums = append(nums, 1, 3)
	s.Show("all:", nums)

	nums = append(nums, 2)
	s.Show("free capacity:", nums)

	nums = append(nums, 4)
	s.Show("no allocation", nums)

	nums = append(nums, 10)
	s.Show("more allocation", nums)
}

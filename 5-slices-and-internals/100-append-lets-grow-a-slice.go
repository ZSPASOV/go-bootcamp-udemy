package main

import "fmt"

func main() {
	// it is a builtin, no need to import
	nums := []int{1, 2, 3}

	// append can't change the passed slice, it returns a new slice
	nums = append(nums, 5)
	fmt.Println(nums)

	// append multiple elements
	nums = append(nums, 7, 10, 15)
	fmt.Println(nums)

	// append one slice to another using the ellipsis operator
	tens := []int{12, 18}
	nums = append(nums, tens...)
	fmt.Println(nums)
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	// declare an uninitialized nil slice
	var uniques []int

loop:
	// you can still use the len function on a nil slice
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		// append function can add new elements to a slice
		uniques = append(uniques, n)

		// a slice doesn't contain any elements right from the start
		// uniques[found] = n
		// found++
	}

	fmt.Println("\n\nuniques:", uniques)
	fmt.Println("\nlength of uniques:", len(uniques))

	sort.Ints(uniques)
	fmt.Println("\nsorted:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	// in the example below sort.Ints function accepts only slices and does not work with arrays
	// we convert the nums array to slice with the [:] syntax
	sort.Ints(nums[:])
	fmt.Println("\nnums:", nums)
}

package main

import "fmt"

func main() {
	// when we use a slice literal to declare a slice, go creates a backing array
	ages := []int{35, 15, 25}
	// when you use a slice literal directly the capacity and the length of the slice be the same
	fmt.Println(len(ages))      // 3
	fmt.Println(cap(ages))      // 3
	fmt.Println(len(ages[0:0])) // 0
	fmt.Println(cap(ages[0:0])) // 3
	ages = ages[0:0]
	fmt.Println(ages) // []
	// now we can reslice back to its full capacity, because there was enough capacity in the backing array :
	fmt.Println(ages[0:3]) // [35 15 25]
	ages = ages[1:3]
	fmt.Println(ages) // [15 25]

	// the capacity is 2 now, and the previous elements can't be used anymore
	fmt.Println(cap(ages)) // 2
}

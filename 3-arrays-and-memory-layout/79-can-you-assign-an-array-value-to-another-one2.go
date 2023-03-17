package main

import "fmt"

func main() {
	prev := [3]string{
		"one",
		"two",
		"three"}

	newer := prev

	for i := range prev {
		newer[i] += " 2nd Ed."
	}

	// notice that the prev array has not changed, only the newer array changed
	fmt.Printf("%#v\n", prev)  // [one two three]
	fmt.Printf("%#v\n", newer) // [one 2nd Ed. two 2nd Ed. three 2nd Ed.]

	var newest [4]string

	// now we cannot copy array prev to array newest, because they are of different length
	// we can add the values of array prev to array newest with a loop
	for i, b := range prev {
		newest[i] += b + " 3rd Ed."
	}
	fmt.Printf("%#v\n", newest)

	newest[3] = "four"
	fmt.Printf("%#v\n", newest)
	newarr := [...]string{"asd", "new"}
	fmt.Println(newarr)

}

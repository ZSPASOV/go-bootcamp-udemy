package main

import (
	"fmt"
)

func main() {
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("[english word] -> [turkish word]")
	// 	return
	// }
	// query := args[0]

	// #1: Nil Map: Read-Only
	var dict map[string]string

	// #5: You cannot assign to a nil map.
	// dict["up"] = "nagore"
	// dict["down"] = "nadolu"

	// #2: Map retrieval is O(1) — on average.
	key := "good"

	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)

	// #1B
	fmt.Printf("# of Keys: %d\n", len(dict))

	// fmt.Printf("Zero Value: %#v\n", dict)

	// #4: Nil map ready to use
	// if dict != nil {
	// 	value := dict[key]
	// 	fmt.Printf("%q means %#v\n", key, value)
	// }

	// #3: Cannot use non-comparable types as map key types
	// var broken map[[]int]int
	// var broken map[map[int]string]bool
	//
	// A map can only be compared to nil value
	// _ = dict == nil
}

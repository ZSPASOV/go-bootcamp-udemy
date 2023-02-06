package main

import "fmt"

func main() {
	var brand = "Google"
	// "%q\n" and "%s\n" are formatting text. determines what and how to print
	// brand in that case is replacer value. it replaces the verbs inside the formatting text 
	// v1 
	// fmt.Printf("%q\n", brand)
	// v2 
	fmt.Printf("%s\n", brand)
}

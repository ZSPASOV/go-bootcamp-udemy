package main

import "fmt"

func main() {
	type person struct {
		name, lastname string
		age            int
	}

	var picasso person
	var freud person

	fmt.Printf("\nPicasso: %+v\n", picasso)
	fmt.Printf("\nFreud: %+v\n", freud)

	picasso.name = "Pablo"
	picasso.lastname = "Picasso"
	picasso.age = 91

	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\nPicasso: %#v\n", picasso)
	fmt.Printf("\nFreud: %#v\n", freud)

	// using a Struct literal

	zapryan := person{name: "Zapryan", lastname: "Spasov", age: 33}
	fmt.Printf("\nFreud: %#v\n", zapryan)
}

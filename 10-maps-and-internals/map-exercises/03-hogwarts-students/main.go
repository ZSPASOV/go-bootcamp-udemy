package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house := os.Args[1]

	houses := map[string][]string{
		"gryffindor": {
			"weasley",
			"hagrid",
			"dumbledore",
			"lupin",
		},
		"hufflepuf": {
			"wenlock",
			"scamander",
			"helga",
			"diggory",
		},
		"ravenclaw": {
			"wenlock",
			"scamander",
			"helga",
			"diggory",
		},
		"slytherin": {
			"horace",
			"nigellus",
			"higgs",
			"scorpius",
		},
		"bobo": {
			"wizardry",
			"unwanted",
		},
	}

	delete(houses, "bobo")

	students, ok := houses[house]
	if !ok {
		fmt.Printf("Sorry. I don't know anything about %#v\n.", house)
	} else {
		fmt.Printf("~~~%#v students~~~\n", house)
		for _, student := range students {
			fmt.Println(student)
		}
	}
}

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	godOfWar := game{
		item: item{
			id:    1,
			name:  "god of war",
			price: 50,
		},
		genre: "action adventure",
	}

	xCom2 := game{
		item: item{
			id:    2,
			name:  "x-com 2",
			price: 30,
		},
		genre: "strategy",
	}

	mineCraft := game{
		item: item{
			id:    3,
			name:  "minecarft",
			price: 20,
		},
		genre: "sandbox",
	}

	games := []game{godOfWar, xCom2, mineCraft}

	gamesDirect := []game{
		game{
			item: item{
				id:    5,
				name:  "Resident Evil 4",
				price: 80,
			},
			genre: "survivor",
		},
		game{
			item: item{
				id:    6,
				name:  "Final Fantasy 10",
				price: 70,
			},
			genre: "rpg",
		},
	}

	fmt.Printf("the games: %#v\n\n", games)
	fmt.Printf("the games direct: %#v\n\n", gamesDirect)

	fmt.Println(games)
	fmt.Println(gamesDirect)

}

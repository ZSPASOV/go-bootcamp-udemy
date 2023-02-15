package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gimme a month name.")
		return
	}

	// v1 - if statement
	// m := os.Args[1]

	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// 	fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Println("ERROR: not a month.")
	// }

	// v2 - short switch statement
	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}

	// switch statement
	m := os.Args[1]
	switch m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	const (
		userName = "panyo"
		password = "asdfgh"
	)
	inputName := os.Args[1]
	inputPass := os.Args[2]
	if len(os.Args) < 2 {
		fmt.Println("Usage: [username] [password]")
	} else if inputName != userName {
		fmt.Printf("Access denied for %v\n.", inputName)
	} else if inputPass != password {
		fmt.Printf("Invalid password for %v\n.", inputPass)
	} else if inputName == userName && inputPass == password {
		fmt.Printf("Access granted to %v.\n", inputName)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("1st argument:", os.Args[3])

	fmt.Println("Number of items inside os.Args:", len(os.Args))
}

// Println only prints the value, with Printf we can get type and value
// go run creates temporary executable files, so every time we run it, different results will be printed
// go build creates permanent executable files
// in the terminal type go build -o greeter. -o allows to give a name to the compiled program
// run ./greeter hi hello hey in the terminal
// run go run 09-the-basics-of-os-args.go  hi yo asd

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Println(i, "--", v)
	}
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Println(i, "--", v)
	}
	for i, v := range os.Args[1:] {
		fmt.Println(i, "--", v)
	}
	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}
}

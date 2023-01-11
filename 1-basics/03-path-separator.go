package main

import (
	"fmt"
	"path"
)

func main() {
	// the split function returns two strings, the first is _ because it is not used
	// var file string
	// _, file = path.Split("css/main.css")
	// fmt.Println(file)
	// v2
	var dir string
	dir, _ = path.Split("css/main.css")
	fmt.Println(dir)
}

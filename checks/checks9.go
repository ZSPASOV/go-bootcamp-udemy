package main

import "fmt"

func main() {
	lyric := []string{"le", "vent", "nous", "portera"}
	n := copy(lyric, make([]string, 4))

	fmt.Printf("%d %q\n", n, lyric)

	// -- USEFUL INFORMATION (but not required to solve the question) --
	//      In the following `copy` operation, `make` won't allocate
	//      a slice with a new backing array up to 32768 bytes
	//      (one string value is 8 bytes on a 64-bit machine).
	//
	//      This is an optimization made by the Go compiler.
}

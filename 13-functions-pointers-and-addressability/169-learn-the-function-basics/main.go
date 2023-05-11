// You need run this program like so:
//   go run .
//
// This will magically pass all the go files in the current directory to the
// Go compiler.
//
//
// BUT NOT like so:
//   go run main.go
//
// Because, the compiler needs to see bad.go too
// It can't magically find bad.go — what you give is what you get.
//

package main

// N is a shared counter which is BAD, declaring package level variables is a bad practice. Every file in the same package can change them.
var N int

func main() {
	showN()
	incrN()
	incrN()
	showN()
	incrN()
	incrN()
	incrN()
	incrN()
	showN()
}

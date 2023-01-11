package main

func main() {

}

// OS PACKAGE
// os = operatig system. Allows you to  access operating system functionalities
// intro to slices:
// Args variable belongs to the os package
// var Args []string - Args` type is a slice of strings
// When you run a Go program, go puts the command-line arguments into this variable automatically
// A slice itself is a single value, it just points to the other values that it stores behind the scenes.
// Each value inside a slice is an unnamed variable
// go run main.go hi yo - hi and yo are arguments, when we use space between them they became two arguments.
// they are stored in the Args slice as string values. They can be accessed by Args[0]  Args[1] Args[2]etc
// Args[0] stores the path to the current executing program.
// Args[1] - first argument - 'hi', Args[2] - second argument 'yo'


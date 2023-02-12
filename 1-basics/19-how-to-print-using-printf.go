package main

import "fmt"

func main() {
	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	// in Printf indexes start at 1, not 0
	fmt.Printf("%v is %v away. Thik! %[2]v kms! %[1]v OMG!\n", planet, distance)
}

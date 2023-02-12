package main

import "fmt"

func main() {
	speed := 100 // speed is int
	force := 2.5 // force is float 64
	// there will be compiler error below, you can`t use values belonging to different types together:
	// speed = speed * force
	// we should convert one of variables to the type of the other, we should convert speed to float
	// because if we convert force to int 2.5 will become 2 
	var new_speed float64
	new_speed = float64(speed) * force
	fmt.Println(new_speed)
	// below first we convert speed to float64 to complete the * operation, then we convert the result to int, because
	// speed is of type int and we should assign only int values to it
	speed = int(float64(speed) * force)
	fmt.Println(speed)

}

// type conversion - converts a value to another type 
// type(value)   type - changes the type of a given value into this type, value - value to be converted
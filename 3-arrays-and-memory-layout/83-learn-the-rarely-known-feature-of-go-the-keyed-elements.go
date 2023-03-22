package main

import "fmt"

func main() {
	// the keyed elements can be in any order
	rates := [3]float64{
		1: 2.5,
		0: 0.5,
		2: 1.5,
	}
	fmt.Printf("%#v\n", rates) // [3]float64{0.5, 2.5, 1.5}

	// uninitialized elements will be initialized to their zero values
	ratesTwo := [3]float64{
		2: 1.5,
	}
	fmt.Printf("%#v\n", ratesTwo) // [3]float64{0, 0, 1.5}

	ratesThree := [3]float64{
		1: 1.5,
	}
	fmt.Printf("%#v\n", ratesThree) // [3]float64{0, 1.5, 0}

	ratesFour := [...]float64{
		5: 1.5,
	}
	fmt.Printf("%#v\n", ratesFour) // [6]float64{0, 0, 0, 0, 0, 1.5}

	// this will not be used in practice, just showing :
	ratesFive := [...]float64{
		5: 1.5,
		2.5,
		0: 0.5,
	}
	fmt.Printf("%#v\n", ratesFive) // [7]float64{0.5, 0, 0, 0, 0, 1.5, 2.5}
}

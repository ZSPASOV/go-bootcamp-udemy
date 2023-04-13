package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	data := []float64{10, 25, 30, 50}

	copy(data, []float64{99, 100})

	s.Show("Data: ", data)

	dataTwo := []float64{37, 45}

	copy(dataTwo, []float64{70, 80, 90, 120, 65})

	s.Show("dataTwo: ", dataTwo)
}

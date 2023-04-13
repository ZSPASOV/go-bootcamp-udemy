package main

import "fmt"

func main() {
	multiDimensionalSlice := [][]int{
		{20, 40, 50, 44},
		{35, 45, 55},
		{70, 90, 100, 400, 550},
	}

	for i, innerSlice := range multiDimensionalSlice {
		for _, value := range innerSlice {
			{
				fmt.Printf("On index : %d --> the value : %d \n", i, value)
			}
		}
	}
}

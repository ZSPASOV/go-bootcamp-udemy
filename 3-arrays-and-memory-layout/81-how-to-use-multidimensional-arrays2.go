package main

import "fmt"

func main() {
	students := [2][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64
	sum += students[0][0] + students[0][1] + students[0][2]
	sum += students[1][0] + students[1][1] + students[1][2]

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avg Grade: %v\n", sum/N)

	// v2 better
	studentsTwo := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sumTwo float64
	for _, grades := range studentsTwo {
		for _, grade := range grades {
			sumTwo += grade
		}
	}
	fmt.Printf("Avg Grade: %v\n", sumTwo/N)
}

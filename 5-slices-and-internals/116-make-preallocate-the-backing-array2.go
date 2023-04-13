package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	var upTasks []string
	s.Show("upTasks", upTasks)

	// v1
	// using that approach always the a new backing array is allocated, an expensive operation
	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
	// v2
	// here a new backing array is not created, but the values are appended at the end of the slice. Also because there was not enough space a new backing array was created, and we want to avoid that.
	upTasksTwo := make([]string, len(tasks))
	s.Show("upTasksTwo", upTasksTwo)
	for _, task := range tasks {
		upTasksTwo = append(upTasksTwo, strings.ToUpper(task))
		s.Show("upTasksTwo", upTasksTwo)
	}

	// v3 
	// now the append function appending starts from the beginning of the slice. that is because the  append function appends at the first place after the length. the slice starts with length 0 in that case.
	upTasksThree := make([]string, 0, len(tasks))
	s.Show("upTasksThree", upTasksThree)

	for _, task := range tasks {
		upTasksThree = append(upTasksThree, strings.ToUpper(task))
		s.Show("upTasksThree", upTasksThree)
	}
}

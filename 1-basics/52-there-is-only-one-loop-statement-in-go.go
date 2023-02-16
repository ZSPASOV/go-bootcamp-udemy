package main

import "fmt"

func main() {
	var sum int
	var sum2 int

	// i := 1 is the INIT STATEMENT. it is a simple statement. it runs once. before the loop begins
	// i <= 5 is the CONDITION EXPRESSION of the for loop. It should be a bool expression only. As long as this condition yields true, the loop will continue. when it becomes false the loop will end
	// i ++ is the POST STATEMENT of the for loop. like the INIT STATEMENT it is a simple statement. This statement will be run after each statement in the loop.
	// ; are SEPARATORS. They separate the parts of a for statement
	// the INIT STATEMENT, THE CONDITION EXPRESSION and the POST statement are all optional

	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	// the INIT STATEMENT, THE CONDITION EXPRESSION and the POST statement are all optional:
	j := 1
	for j <= 5 {
		sum2 += j
		j++
	}

	fmt.Println(sum2)
}

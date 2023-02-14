package main

func main() {
	//similar to short if statement
	// you should use a semicolon to separate the simple statement and the case conditions
	// the ; separator separetes the simple statement and the switch condition
	switch i := 10; {
	case i > 0:
		// "positive"
	case i < 0:
		// "negative"
	default:
		// "zero"
	}
}

package main

import "fmt"

func main() {

	dictionary := map[string]string{
		"good":      "dobre",
		"very good": "mnogo dobre",
		"perfect":   "perfektno",
	}

	// printira podredeno po azbuchen red
	fmt.Printf("%#v\n", dictionary)
}

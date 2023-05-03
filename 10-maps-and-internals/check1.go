package main

import "fmt"

func main() {

	dictionary := map[string]string{
		"good":      "dobre",
		"very good": "mnogo dobre",
		"perfect":   "perfektno",
	}

	for key, value := range dictionary {
		fmt.Println("key: ", key, ", value:", value)
	}

	dictionary["bad"] = "zle"
	dictionary["perfect"] = "super dobre"
	fmt.Println("overwritten:")

	for key, value := range dictionary {
		fmt.Println("key: ", key, ", value:", value)
	}

	fmt.Println(dictionary["very good"])
	fmt.Println("ne e namereno:", dictionary["not found"])

	if dictionary["not found"] == "" {
		fmt.Println("nishto ne nameri")
	}

	// v2 - using ok

	value, ok := dictionary["not found"]
	fmt.Println(value)
	fmt.Println(ok)
}

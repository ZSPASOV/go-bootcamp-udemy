package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [yearly]string{"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}

	books[3] = "Kafka's Revenge 2nd Edition"

	fmt.Printf("books : %#v\n", books)
}

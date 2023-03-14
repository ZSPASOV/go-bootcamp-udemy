package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]

	// v 1
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	// v2
	for i := 0; i < len(sBooks); i++ {
		sBooks[i] = books[i+1]
	}

	// v3
	for _, v := range sBooks {
		fmt.Println(v)
		// v is a copy, it is not the original element, so the original values in sBooks won't be affected
		v += "won't effect"

	}

	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	// in that case we are passing the number of elements of books as number of elements to the new array published.
	// it will contain only bool elements
	var published [len(books)]bool
	fmt.Println(published)

	published[0] = true
	published[len(books)-1] = true
	fmt.Println("\nPublished Books:")

	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}

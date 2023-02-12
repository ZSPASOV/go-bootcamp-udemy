package main

// there is no subpackaging in go. utf8 is in the unicode folder, but it is not a subpackage of unicode
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// len built-in function returns the length of a string value in bytes
	name := "carl"
	fmt.Println(len(name)) // 4 - it returns the number of characters for English letters only
	jp := "ゴシック"
	fmt.Println(len(jp)) // 12 - for Japanese letters in that case it returns the number of bytes
	// Unicode characters can be 1 to 4 bytes each

	// utf8 package's RuneCountInString function finds the number of characters in a string
	fmt.Println(utf8.RuneCountInString(jp)) // 4

}

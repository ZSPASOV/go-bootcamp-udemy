package main

import "fmt"

func main() {
	var s string
	s = "how are you" // string literal
	s = `how are you` // raw string literal
	fmt.Println(s)
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>" // string literal (this is hard to read and maintain, it is better to use raw string literal instead, so not need to write escape sequences)
	fmt.Println(s)
	s = `
<html>
	<body>"Hello"</body>
</html>` // raw string literal
	fmt.Println(s)
	fmt.Println("c:\\my\\dir\\file") // string literal
	fmt.Println(`c:\my\dir\file`)    // raw string literal
}

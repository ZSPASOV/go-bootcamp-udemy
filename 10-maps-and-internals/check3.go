package main

import "fmt"

func main() {

	multiplePhoneNums := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	phone, ok := multiplePhoneNums["spasov"]
	fmt.Println(phone)
	fmt.Println(ok)

	phone, ok = multiplePhoneNums["dulin"]
	fmt.Println(phone)
	fmt.Println(ok)

	if ok {
		fmt.Println(phone)
	}
}

package main

import "fmt"

func main() {
	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}

	fmt.Printf("1 BTC is %v ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %v WAN\n", rates[WAN])
	fmt.Printf("%#v\n", rates)
}

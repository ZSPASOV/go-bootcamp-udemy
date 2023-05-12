package main

import "fmt"

func main() {
	stats := map[int]int{1: 10, 10: 2}
	incrAll(stats)
	fmt.Println(stats)
	num := 20
	incrNum(num)
	fmt.Println(num)
}

func incrAll(stats map[int]int) {
	for k := range stats {
		stats[k]++
	}
}

func incrNum(num int) {
	num++
}

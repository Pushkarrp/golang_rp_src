package main

import (
	"fmt"
)

func main() {
	fmt.Println("Adding Elements to a Map : ")

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	//adding new elements
	m["four"] = 4
	m["eleven"] = 11

	fmt.Println(m)
	fmt.Println()

	fmt.Println("Looping over a Map : ")

	for k, v := range m {
		fmt.Println(k, v)
	}
}

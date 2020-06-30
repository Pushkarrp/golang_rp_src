package main

import (
	"fmt"
)

func main() {
	x := []int{2, 4, 6, 8, 10}
	//COMPOSITE LITERAL -> type {values}
	// []int {2,4,6,8}
	//SLICE ->  holds VALUES of the same TYPE
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println()

	fmt.Println("We can loop over the values in a slice with the range clause.")

	for i, v := range x {
		// for INDEX, VALUE iterated over RANGE `x` (slice varb)
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println("We can also access items in a slice by index position.")
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

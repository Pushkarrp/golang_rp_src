package main

import (
	"fmt"
)

func main() {
	x := []int{5, 10, 15, 20, 25, 30, 35, 40}
	//COMPOSITE LITERAL => type{values}
	fmt.Println("Cut Parts of the slice away. Use of COLON OPERATOR ")
	//fmt.Println()

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println()

	fmt.Println("Idtfr [ start idx : end idx ] :-")
	fmt.Println("---------------------------------")
	fmt.Print("\nprint all values :\n\t", x[:])
	fmt.Println("\nprint values from zero index to defined-1 index :\n\t", x[:5])
	fmt.Println("\nprint values from defined index to final :\n\t", x[4:])
	fmt.Println("\nprints values from index 1 (incl) to index 5 :\n\t", x[1:6])
}

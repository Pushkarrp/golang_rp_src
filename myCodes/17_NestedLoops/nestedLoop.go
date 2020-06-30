package main

import (
	"fmt"
)

func main() {
	fmt.Println("NESTED FOR LOOP : ")
	fmt.Println("EXAMPLE 1 :")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("EXAMPLE 2 :")
	for i := 0; i < 3; i++ {
		fmt.Printf("Outer %d \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t Inner %d \n", j)
		}
	}

}

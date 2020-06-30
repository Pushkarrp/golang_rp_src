package main

import (
	"fmt"
)

func main() {
	// Slice : [ empty ] Type {values}
	slx := []int{3, 6, 9, 12, 15, 18, 21, 24}

	for i, v := range slx {
		fmt.Println(i, v)
	}

	fmt.Printf("%T \n", slx)
}

package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 4, 6, 8, 10, 12}
	// []int{values} is Composite Literal

	fmt.Println("Loop / Range over a Slice : ")
	fmt.Println("----------------------------")

	//for index , value := RANGE over Slice named `xi`
	for i, v := range xi {
		fmt.Println(i, v)
	}
}

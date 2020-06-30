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
}

package main

import (
	"fmt"
)

func main() {
	// Array : [length specified] Type {values}
	ar := [5]int{5, 10, 15, 20, 25}

	for i, v := range ar {
		fmt.Println(i, v)
	}

	fmt.Printf("%T \n", ar)
}

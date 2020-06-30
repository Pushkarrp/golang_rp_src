package main

import (
	"fmt"
)

func main() {
	sx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	/*
		for i,v := range sx {
			fmt.Println(i,v)
		}
	*/

	fmt.Println(sx[0:5])
	fmt.Println(sx[5:])
	fmt.Println(sx[2:7])
	fmt.Println(sx[1:6])
}

package main

import (
	"fmt"
)

func main() {
	sx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("Deletion in Middle :")
	fmt.Println("--------------------")
	sx = append(sx[:3], sx[6:]...)
	fmt.Println(sx)

	fmt.Println()
	fmt.Println("Deletion from Start :")
	fmt.Println("----------------------")
	sx = sx[4:]
	fmt.Println(sx)

}

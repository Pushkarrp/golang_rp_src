package main

import (
	"fmt"
)

func main() {
	sx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Append single element
	sx = append(sx, 52)
	fmt.Println(sx)

	// Append multiple independent elements
	sx = append(sx, 53, 54, 55)
	fmt.Println(sx)

	y := []int{56, 57, 58, 59, 60}

	// Append another Slice `...`
	sx = append(sx, y...)
	fmt.Println(sx)

}

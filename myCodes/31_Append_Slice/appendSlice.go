package main

import (
	"fmt"
)

func main() {
	fmt.Println("Append = Adding at End")
	fmt.Println("Use : BUILT-IN func called : `append()` ")
	fmt.Println()
	x := []int{1, 2, 3}
	fmt.Println("Before : x = ", x)

	//appending VALUES at end
	x = append(x, 4, 5, 6, 7)
	fmt.Println("After : x = ", x)

	y := []int{9, 10, 11, 12}
	//appending another SLICE at end
	// `...` Represents => Another Group of Values. NOT Individual Values
	x = append(x, y...)
	fmt.Println("After after : x = ", x)

}

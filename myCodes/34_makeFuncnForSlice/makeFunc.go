package main

import (
	"fmt"
)

func main() {
	fmt.Println(" We can use the built-in function 'make()' to specify :")
	fmt.Println("1. How large our slice should be &")
	fmt.Println("2. Also how large the underlying array should be.")
	fmt.Println("This can enhance performance a little bit.")

	// make( []TYPE , LENGTH of Slice , CAPACITY of Underling array)
	x := make([]int, 4, 7)

	fmt.Println("slice x : ", x)
	fmt.Println("len : ", len(x))
	fmt.Println("cap : ", cap(x))

	for i := 0; i <= 2; i++ {
		x[i] = (i + 1) * 10
	}

	fmt.Println("slice x : ", x)
	fmt.Println("len : ", len(x))
	fmt.Println("cap : ", cap(x))

	x = append(x, 50) //Adds element at index 5
	fmt.Println("slice x : ", x)
	fmt.Println("len : ", len(x))
	fmt.Println("cap : ", cap(x))

	x = append(x, 60, 70, 80)
	fmt.Println("slice x : ", x)
	fmt.Println("len : ", len(x))
	fmt.Println("cap : ", cap(x))
	fmt.Println("Capacity of 7 gets overFlown i.e array size goves above len 7")
	fmt.Println("So new Capacity of Underlying array becomes double i.e 14")

}

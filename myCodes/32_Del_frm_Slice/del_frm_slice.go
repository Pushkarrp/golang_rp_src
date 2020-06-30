package main

import (
	"fmt"
)

func main() {
	fmt.Println("DELETING FROM A SLICE : ")
	fmt.Println("----------------------")
	fmt.Println("SIMPLE TARIKA : Use `append()` func for deleting elements")
	fmt.Println("----------------------------------------------------------")

	x := []int{10, 20, 30}

	// individual values. So n `...` used
	x = append(x, 40, 50, 60, 70, 80)

	fmt.Println("Before Deletion : ", x)

	fmt.Println("we need to delete values 30,40,50 from Slice x : ")

	//ASLI DELETION BABYY
	// `...` represents -> another Aggregate dataType
	x = append(x[:2], x[5:]...)

	fmt.Println("After Deletion : ", x)

}

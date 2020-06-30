package main

import (
	"fmt"
)

func main() {
	fmt.Println("FOR LOOP which is kinda While-ish : ")
	fmt.Println("For with only a Single condition")

	// init ouside
	// for condition {  post inside }

	x := 1
	for x <= 10 {
		fmt.Print(x*2, "  ")
		x++
	}
}

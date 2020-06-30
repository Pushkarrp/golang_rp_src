package main

import (
	"fmt"
)

func main() {

	fmt.Println("If-Else : ")
	x := 42
	if x == 100 {
		fmt.Println("value of x is 100")
	} else {
		fmt.Println("value of x is Not 100")
	}

	fmt.Println("If- Else If - Else : ")
	clr := "Yellow"

	if clr == "Red" {
		fmt.Println("STOP")

	} else if clr == "Yellow" {
		fmt.Println("WAIT")

	} else if clr == "Green" {
		fmt.Println("Go")

	} else {
		fmt.Println(" :/ ")

	}

}

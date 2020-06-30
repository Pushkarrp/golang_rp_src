package main

import (
	"fmt"
)

func main() {
	clr := "Green"

	if clr == "Red" {
		fmt.Println("Red = Stop")
	} else if clr == "Yellow" {
		fmt.Println("Yellow = Wait")
	} else if clr == "Green" {
		fmt.Println("Green = Go")
	} else {

	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch keyword with no value after it means => switch true")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println()

	switch { //switch true
	case false:
		fmt.Println("ain't gona print baby")
	case 42 == 21:
		fmt.Println("!Print")
	case true:
		fmt.Println("Satyamev Jayate")
	}
}

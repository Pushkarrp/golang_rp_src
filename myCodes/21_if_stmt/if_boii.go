package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	fmt.Println("If with INITIALIZATION Statement : ")
	// x initalized ; condition checked
	if x := 42; x == 24 {
		fmt.Println("Yo in x")
	}
}

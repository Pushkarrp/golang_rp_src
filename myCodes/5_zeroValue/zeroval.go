package main

import (
	"fmt"
)

var y string

// since no value assigned varb y will hold ZERO VALUE for Strings

var n int

// since no value assigned varb n will hold ZERO VALUE for Integers

func main() {
	fmt.Println("value of y : ", y, "||")
	// y prints an EMPTY STRING

	y = "JAH JAH"
	fmt.Println("value of y : ", y, "||")

	fmt.Println("value of n : ", n, "||")

	/*
		ZERO VALUES
		============
		false for booleans

		0 for integers

		0.0 for floats

		"" for strings

		nil for
			pointers
			functions
			interfaces
			slices
			channels
			maps
	*/
}

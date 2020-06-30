package main

import (
	"fmt"
)

// VAR can be used to declare variables outside a function. i.e Allows Global scope
var z = 900

var num int

// DECLARE there is a VARIABLE with the IDENTIFIER "num"
// and that the VARIABLE with the IDENTIFIER "num" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "num"

// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps

func main() {
	x := 200
	// S.D.O can only be used to declr a varb inside Func. i.e Local scope
	fmt.Println("x (by SDO)= ", x)

	var y = 400
	fmt.Println("y (by var)= ", y)

	fmt.Println("z (by var outside)= ", z)

	foo()
} //main() ends

func foo() {
	fmt.Println("again z = ", z)
}

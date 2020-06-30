package main

import (
	"fmt"
)

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"
var z string = "Shaken, not stirred"

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language like python or js

var k string = ` James said, 
"Shaken, 
But
not stirred" `

func main() {
	fmt.Println(y)
	fmt.Printf("%T \n", y)

	fmt.Println(z)
	fmt.Printf("%T \n", z)

	fmt.Println(k)
	//fmt.Printf("%T \n",k)
}

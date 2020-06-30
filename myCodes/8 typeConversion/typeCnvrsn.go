package main

import (
	"fmt"
)

var a int = 42

type bhindi int

// user defd type bhindi with underline build in type as int
var k bhindi = 69

func main() {
	fmt.Print(a, "\t")
	fmt.Printf("%T \n", a)

	fmt.Print(k, "\t")
	fmt.Printf("%T \n", k)

	fmt.Println("TYPE CONVERSION :=")
	//TYPE CONVERSION (GOlang name for Type Casting)
	a = int(k) //value of varb k gets assigned to varb a

	fmt.Print(a, "\t")
	fmt.Printf("%T \n", a)

}

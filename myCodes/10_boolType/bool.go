package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value of bool is FALSE
	x = true       // already declared so no SDO:=
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a != b)
}

package main

import "fmt"

type kanda int

var x kanda

func main() {
	fmt.Println(x) //prints zero value
	fmt.Printf("%T \n", x)

	x = 42
	//declared already so only = used and not :=
	fmt.Println(x)

}

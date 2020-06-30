package main

import "fmt"

type kanda int

var x kanda
var y int

func main() {
	fmt.Println(x) //prints zero value
	fmt.Printf("%T \n", x)

	x = 42
	//declared already so only = used and not :=
	fmt.Println(x)

	//==============================================
	fmt.Println("==========================")
	y = int(x)
	// CONVERTS varb x (of type KANDA) into type INT
	//the after Conversion assigns value of varb x to varb y

	fmt.Println(y)
	fmt.Printf("%T \n", y)
}

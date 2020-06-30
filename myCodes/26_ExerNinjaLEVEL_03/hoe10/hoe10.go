package main

import (
	"fmt"
)

func main() {

	//be1 := true && true
	be2 := true && false
	//be3 := true || true
	be4 := true || false
	be5 := !true

	//mt.Printf("true && true : %v\n\n",be1)
	fmt.Printf("true && false : %v\n\n", be2)
	//fmt.Printf("true || true : %v\n\n",be3)
	fmt.Printf("true || false : %v\n\n", be4)
	fmt.Printf("!true : %v\n", be5)

}

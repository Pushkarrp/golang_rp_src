package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println("y = ", y)
	fmt.Printf("VALUE : %v \n", y)
	fmt.Printf("BINARY : %b \n", y)
	fmt.Printf("HEX : %x \n", y)

	fmt.Printf("HEX w/ 0x : %#x\n", y)
	fmt.Printf("Printf : %#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	//prints to a string
	fmt.Println("Sprintf : ", s)
}

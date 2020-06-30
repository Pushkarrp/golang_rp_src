package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {
		// DEC CODE & UTF-8 CODE & HEX CODE & CHAR/ASCII CODE
		fmt.Printf("%v \t\t %#U \t\t %#X \t\t %c\n", i, i, i, i)

	}

	fmt.Printf("\n\n")

	for x := 97; x <= 122; x++ {
		// DEC CODE & UTF-8 CODE & HEX CODE & CHAR/ASCII CODE
		fmt.Printf("%v \t\t %#U \t\t %#X \t\t %c\n", x, x, x, x)
	}

	//fmt.Println()
}

package main

import (
	"fmt"
)

func main() {
	num := 42

	fmt.Printf("%v \t %b \t %#x \n", num, num, num)

	num = num << 1 //LEFT SHIFT ONE BIT

	fmt.Printf("%v \t %b \t %#x \n", num, num, num)
}

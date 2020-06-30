package main

import (
	"fmt"
)

func main() {
	fmt.Println("5 KA PAHADA :")

	for i := 1; i <= 50; i++ {

		if i%5 == 0 {
			fmt.Print(i, " ")
		}
	}
}

package main

import (
	"fmt"
)

func main() {

	for al := 65; al <= 90; al++ {
		fmt.Println(al)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", al)
		}
		fmt.Println()
	}
}

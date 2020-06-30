package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 41; i++ {
		fmt.Printf("%v MOD 4 = %v\n", i, i%4)
	}
}

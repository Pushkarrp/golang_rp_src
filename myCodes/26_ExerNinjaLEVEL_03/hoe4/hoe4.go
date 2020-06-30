package main

import (
	"fmt"
)

func main() {
	bd := 1997
	// for  {}
	for {

		if bd > 2020 {
			break
		}

		fmt.Print(bd, " ")
		bd++
	}
}

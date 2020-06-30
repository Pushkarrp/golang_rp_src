package main

import (
	"fmt"
)

const (
	// IOTA is SUCCESSIVELY INCREASING Pre-Defined Identifier
	// iota initial value starts with ZERO / 0
	// RESET to ZERO whenever Keyword 'const' appears
	yr1 = 2020 + iota //iota = 0
	yr2 = 2020 + iota //iota = 1
	yr3 = 2020 + iota //iota = 2
	yr4 = 2020 + iota //iota = 3
)

func main() {
	fmt.Println(yr1)
	fmt.Println(yr2)
	fmt.Println(yr3)
	fmt.Println(yr4)
}

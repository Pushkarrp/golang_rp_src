package main

import (
	"fmt"
)

const (

	// IOTA is a SUCCESSIVELY INCREASING Pre-Defined Identifier
	// iota initial value starts with ZERO / 0
	// RESET to ZERO whenever Keyword 'const' appears

	a = iota
	b
	c
)

const (

	// IOTA is a SUCCESSIVELY INCREASING Pre-Defined Identifier
	// iota initial value starts with ZERO / 0
	// RESET to ZERO whenever Keyword 'const' appears

	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println("==================")
	dets := `
	
	IOTA is a SUCCESSIVELY INCREASING Pre-Defined Identifier
	iota initial value starts with ZERO / 0
	Resets to ZERO whenever Keyword 'const' appears
	`
	fmt.Println(dets)
}

package main

import (
	"fmt"
)

const n int = 42

const (
	pi = 3.14
	s  = "jah jah"
)

func main() {
	fmt.Println(n)
	fmt.Printf("%T \n", n)
	fmt.Println(pi)
	fmt.Printf("%T \n", pi)
	fmt.Println(s)
	fmt.Printf("%T \n", s)
}

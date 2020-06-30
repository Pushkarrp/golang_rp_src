package main

import (
	"fmt"
)

func main() {
	a := (42 == 45)
	b := (42 <= 45)
	c := (42 >= 45)
	d := (42 != 45)
	e := (42 < 45)
	f := (42 > 45)

	fmt.Println(a, b, c, d, e, f)
}

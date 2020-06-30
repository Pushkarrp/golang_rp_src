package main

import (
	"fmt"
)

func main() {
	fmt.Println("Intro To Maps : ")
	fmt.Printf("\t Maps are Key-Value Store \n\t Un-ordered List \n")
	fmt.Printf("Map Creation : \n\t map_Varb := map [key_Type] value_type { K : V , }")
	// map [key_Type] value_type { K : V , } ---> Composite Type
	fmt.Println()

	m := map[string]int{
		"key":   271997,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println()

	fmt.Println("Map : ", m)
	fmt.Println("k_one = ", m["one"])
	fmt.Println("k_three = ", m["three"])
	fmt.Println()

	fmt.Println("If the given Key is Not Mapped to any Value :")
	fmt.Println("      then Value printed is 0 ")
	fmt.Println("for eg. :- Key_Forty => ", m["forty"])
	fmt.Println()

	fmt.Println("`Comma-Ok` Idiom : ")
	fmt.Println("----------------")
	// v : value  |  ok : bool type
	//ok = True if key Present ; False if key Absent
	v, ok := m["forty"]
	fmt.Println("value : ", v)
	fmt.Println("OK key_forty : ", ok)
	fmt.Println()

	// if { some INITIALIZATION }   ;   { IDENTIFIER to be Checked if T/F }
	if v, ok := m["eleven"]; ok {
		fmt.Println("Key `eleven` is present. vaule : ", v)
		//Not gonna be printed yo
	}

	// if { some INITIALIZATION }   ;   { IDENTIFIER to be Checked if T/F }
	if v, ok := m["two"]; ok {
		fmt.Println("Key `two` is present. vaule : ", v)
	}

}

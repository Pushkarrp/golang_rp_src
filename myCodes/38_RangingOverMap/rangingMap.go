package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ranging/Looping over a Map")
	fmt.Println("--------------------------")

	mymap := map[string]int{
		"ten":    10,
		"twenty": 20,
		"thirty": 30,
		"forty":  40,
		"fifty":  50,
	}

	fmt.Printf("Key \t Value \n")
	for k, v := range mymap {
		fmt.Printf("%s \t %d\n", k, v)
	}
}

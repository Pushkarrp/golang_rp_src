package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps:")
	fmt.Println("key = weight class")
	fmt.Println("value = current UFC champ")
	fmt.Println("--------------------------")
	fmt.Println()

	//varb := map [keyType] valueType { K : V , }
	ufc := map[string]string{
		"Lightweight":  "Khabib Nurmagomedov",
		"Welterweight": "Kamaru Usman",
		"Middleweight": "Israel Adesanya",
	}

	//fmt.Println(ufc)

	for k, v := range ufc {
		fmt.Println(k, "  :  ", v)
	}
}

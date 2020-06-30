package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map Syntax :  ||  varb := map [keyType] valueType { K : V , }")
	fmt.Println("-------------------------------------------------------------")

	fmt.Println("KEY `string` = weight class")
	fmt.Println("VALUE `string Slice[]` = current UFC champ , country ")
	fmt.Println("-----------------------------------------------------")
	fmt.Println()

	//varb := map [keyType] valueType { K : V , }
	    //[single_String]  String_Slice
 	ufc := map[string][]string{
		"Lightweight":  []string{"Khabib Nurmagomedov", "Russia"},
		"Welterweight": []string{"Kamaru Usman", "Nigeria"},
		"Middleweight": []string{"Israel Adesanya", "Nigeria"},
	}

	/*
	fmt.Println("Simple looping")
	fmt.Println("---------------")
	for k, v := range ufc {
		fmt.Println(k, "  :  ", v)
	}
	fmt.Println()
	fmt.Println()
	*/	


	fmt.Println("Deeper looping")
	fmt.Println("---------------")

	for ky,valAry := range ufc {
		fmt.Println("Weight Class : ",ky)

		for _,dets := range valAry {
			fmt.Printf("\t\t%v \n",dets)
		}
		fmt.Println()
	}
}

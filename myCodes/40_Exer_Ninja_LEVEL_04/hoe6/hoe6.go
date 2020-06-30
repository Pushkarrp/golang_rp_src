package main

import (
	"fmt"
)

func main() {
	fmt.Println("Create Slice using `make()` function :")
	fmt.Println("---------------------------------------")

	fmt.Println("Syntax : slx := make( []Type , Length , Capcty ) ")
	fmt.Println("--------------------------------------------------")

	// create slice
	kl := make([]string, 4, 6)

	// add elements manually
	kl[0] = "section 80"
	kl[1] = "good kid, m.a.a.d city"

	fmt.Println("After manual insertion : ")
	fmt.Println("--------------------------")
	fmt.Println("len : ", len(kl), " ; as assigned")
	fmt.Println("cap : ", cap(kl), " ; as assigned")
	fmt.Println("initial Slice : ", kl)
	fmt.Println()

	// Append Values
	// Values added to Index : length
	//leaving in b/w spaces empty
	kl = append(kl, "good kid, m.a.a.d city", "to pimp a butterfly", "damn")

	fmt.Println("After Appending : ")
	fmt.Printf("\t Values added to index posn : length + 0 \n")
	fmt.Printf("\t leaving in b/w spaces empty \n")
	fmt.Println("--------------------------------------------------------------------------")
	for i, v := range kl {
		fmt.Println(i, v)
	}
	fmt.Println("len : ", len(kl), " ; New Len Assigned")       //New Length Assigned
	fmt.Println("cap : ", cap(kl), " ; new Cap = 2 * old Cap ") // new capacity = 2 * old capacity
}

package main 

import (
	"fmt"
)

func main(){
	fmt.Println("KEY 'string' : Artist")
	fmt.Println("VALUE 'string' : Last Album")

	art := map[string]string{
		"Kenny" : "Damn",
		"Cole" : "KOD",
		"Abel" : "AfterHours",
	}
	fmt.Println()
	fmt.Println(art)

	fmt.Println()
	fmt.Println("Adding a new key-value pair :")
	fmt.Println("varb[`New_key`] = `New_value(s)`")
	fmt.Println("------------------------------")

	art["Tyler"] = "IGOR"

	fmt.Println("Looping Boi :")
	fmt.Println("--------------")
	for ky,val := range art{
		fmt.Println(ky," : ",val)
	}
}
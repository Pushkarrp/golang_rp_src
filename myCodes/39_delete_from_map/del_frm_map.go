package main

import (
	"fmt"
)

func main(){
	fmt.Println("Deletion of Elements from map : ")
	fmt.Println("-------------------------------")

	mymap := map[string] int{
		"two" : 2,
		"four" : 4,
		"six" : 6,
		"eight" : 8,
		"ten" : 10,
	}
	//Loop over map
	for key,val := range mymap{
		fmt.Println(key,val)
	}
	fmt.Println()


	fmt.Println("Asli Deletion Boi : ")
	fmt.Println("--------------------")
	fmt.Printf("General Syntax:\n\t\t `  delete( <mapVarb> , 'key' )  `  \n")
	delete(mymap, "four")
	delete(mymap,"eight")
	delete(mymap,"six")

	fmt.Println()
	fmt.Println("After Deletion Map remaining is : ")
	for k,v :=range mymap{
		fmt.Println(k,v)
	}
}
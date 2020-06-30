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
		"Tyler" : "IGOR",
	}
	fmt.Println("Our Map Boi : ",art)
	fmt.Println()

	fmt.Println("Deletion From Map :")
	fmt.Printf("Syntax || delete(`mapVarb` , 'Del_Key') \n")
	fmt.Println("--------------------------------------")
	
	//delte syntax : delete(`mapVarb` , `Del_Key`) 
	delete(art , "Abel")


	fmt.Println()
	fmt.Println("Looping Boi :")
	fmt.Println("--------------")
	for ky,val := range art{
		fmt.Println(ky," : ",val)
	}
}
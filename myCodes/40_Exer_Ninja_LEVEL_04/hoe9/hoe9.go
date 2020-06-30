package main

import (
	"fmt"
)

func main(){
	fmt.Println("KEY 'String' : Weight Class")
	fmt.Println("VALUE 'String Slice' : Current Champ ,Country, MMA Record")
	fmt.Println("----------------------------------------------")

	ufc := map[string][]string{
		"Lightweight" : []string{"Khabib Nurmagomedov","Russia", "28-0-0"},
		"Welterweight": []string{"Kamaru Usman", "Nigeria","16-1-0"},
		"Middleweight": []string{"Israel Adesanya", "Nigeria","19-0-0"},
	}

	fmt.Println("Adding New Value :")
	fmt.Println(" varb[KEY] = VALUE(s) ")
	fmt.Println("----------------------")
	ufc["Light_Heavyweight"] = []string {"Jon Jones","America","26-1-0"} 

	for ky,valAr := range ufc {
		fmt.Println("Weight Category : ",ky)
		fmt.Println("--------------------------------------")

		for _,detls := range valAr {
			fmt.Println(detls)
		}
		fmt.Println()
	}
	
}
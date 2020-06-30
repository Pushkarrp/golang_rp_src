package main

import (
	"fmt"
)

func main() {
	fmt.Println("Multi Dim Slice")
	fmt.Println("----------------")

	xs1 := []string{"cole", "kendrick", "weeknd"}
	xs2 := []string{"KOD", "Damn", "afterHours"}
	xsfin := [][]string{xs1, xs2}

	fmt.Println(xsfin)
	fmt.Println()

	for i, v := range xsfin {
		fmt.Println(i, v)
	}

	fmt.Println()

	for i, indSx := range xsfin {
		fmt.Println("Row : ", i)

		for j, val := range indSx {
			fmt.Printf("\t index : %v \t value : %v \n", j, val)
		}
	}

}

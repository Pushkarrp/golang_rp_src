package main

import (
	"fmt"
)

func main() {
	favSport := "MMA"

	switch favSport { //switch true
	case "Football":
		fmt.Println("Lionel Messi")
	case "Basketball":
		fmt.Println("Lebron James")
	case "MMA":
		fmt.Println("Khabib Nurmagomedov")
	case "Boxing":
		fmt.Println("Tyson Fury")
	default:
		fmt.Println("IDk man")
	}
}

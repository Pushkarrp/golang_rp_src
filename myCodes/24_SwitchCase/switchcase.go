package main

import (
	"fmt"
)

func main() {

	fmt.Printf("SWITCH CASE\n============\n")

	fmt.Println("EXAMPLE 1 :")
	fmt.Println("Switch on Bool Value :")
	fmt.Println("NOTE : Golang does Not Have FallThrough :")

	switch { //nothing In front of Switch => means TRUE is there
	case false:
		fmt.Println("won't print")

	case 2 != 2:
		fmt.Println("Not print")

	case 5 == 5:
		fmt.Println("Issa True")

	case 9 == 9:
		fmt.Println("TRUE ; does it Print ???")
	}

	fmt.Printf("\n=================\n")
	fmt.Println("Example 2 : ")
	fmt.Println("Explicit Fallthrough by Us : ")

	switch { //nothing In front of Switch => means TRUE is there
	case 5 != 5:
		fmt.Println("no print as False")
	case 2 == 2:
		fmt.Println("True 001")
		fallthrough //explicitly added fallthrugh
	case 3 == 3:
		fmt.Println("True 002")
		fallthrough
	case false:
		fmt.Println("__False__ ; Still Prints cuz FallThrough")
	}

	//=================================================

	fmt.Printf("\n=================\n")
	fmt.Println("Example 3 : ")
	fmt.Println("DEFAULT Case : ")
	switch {
	case 2 != 2:
		fmt.Println("no print")
	case false:
		fmt.Println("No print 2")
	case 5 == 10:
		fmt.Println("No Print 003")
	default:
		fmt.Println("Defalut says => Nothing True ; So I'll Do ;) ")
	}

	//====================================

	fmt.Printf("\n=================\n")
	fmt.Println("Example 4 : ")
	fmt.Println("Switch on a value ; The Case that Matches : ")

	switch "ColeWorld" {
	case "KungFuKenny":
		fmt.Println("Be Humble")
	case "Drizzy":
		fmt.Println("Where the Real Friends at??")
	case "ColeWorld":
		fmt.Println("Ain't Nothin Sunny!!")
	}

	//==========================================
	fmt.Printf("\n=================\n")
	fmt.Println("Example 5 : ")
	fmt.Println("Switch on Multiple Matches for a Case : ")

	num := 10

	switch num {
	case 10, 11, 12:
		fmt.Println("10 or 11 or 12")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("Default")
	}

} //main ends

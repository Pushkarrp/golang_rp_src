package main

import "fmt"

func main() {
	fmt.Println("Hello Homies. Issa Pushkar.")

	foo()

	for i := 1; i <= 11; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Im in foo.")
}

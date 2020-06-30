package main

import "fmt"

//Strings are IM-MUTABLE in Golang

func main() {
	s := "James Bond"
	drk := `This is done with
	BACK-QUOTES ;) `
	// Them ni99as `...` back-quotes Above the Tab button
	fmt.Println(s)
	fmt.Println(drk)

	wrd := "Look at my ASCII"
	fmt.Println(wrd)

	bs := []byte(wrd)
	fmt.Println(bs)
	fmt.Printf("%T \n\n", bs) //uint8 = Byte

	//UTF-8 Character points
	// each code point is called a rune = uint32
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U  ", s[i])
	}
	fmt.Println()

} //main()

package main

import (
	"fmt"
)

func main() {
	fmt.Println("2D Slice : ")

	jc := []string{"J", "Cole", "2014FHD"}
	kl := []string{"Kendrick", "Lamar", "TPAB"}

	fmt.Println(jc)
	fmt.Println(kl)

	multiDim := [][]string{jc, kl}
	fmt.Println("Multi Dim Slice : ", multiDim)

}

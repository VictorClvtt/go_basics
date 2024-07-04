package main

import "fmt"

func main() {

	x := []int{123, 321, 213, 312, 132, 231, 111, 222, 333, 131}

	// A primeira variável recebe os índices e a segunda recebe os valores correspondentes
	// for que exibe do primeiro ao terceiro item do slice
	for i, v := range x[:3] {
		fmt.Println(i, ":", v)
	}
	fmt.Println("----------")
	// for que exibe do quinto ao ultimo item do slice
	for i, v := range x[4:] {
		fmt.Println(i, ":", v)
	}
	fmt.Println("----------")
	// for que exibe do segundo ao setimo item do slice
	for i, v := range x[1:7] {
		fmt.Println(i, ":", v)
	}
	fmt.Println("----------")
	// for que exibe do terceiro ao penultimo item do slice
	for i, v := range x[2 : len(x)-1] {
		fmt.Println(i, ":", v)
	}
}

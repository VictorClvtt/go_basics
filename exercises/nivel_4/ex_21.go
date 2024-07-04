package main

import "fmt"

func main() {

	x := []int{123, 321, 213, 312, 132, 231, 111, 222, 333, 131}

	// A primeira variável recebe os índices e a segunda recebe os valores correspondentes
	for i, v := range x {
		fmt.Println(i, ":", v)
	}
	fmt.Printf("%T", x)
}

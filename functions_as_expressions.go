package main

import "fmt"

func main() {

	num := 367

	x := func(num int) {
		fmt.Print(num, " vezes 2 é: ")
		fmt.Println(num * 2)
	}

	x(num)

}

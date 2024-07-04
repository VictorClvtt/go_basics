package main

import "fmt"

func main() {

	// primeiro closure
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// segundo closure
	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func i() func() int {

	var x int

	return func() int {
		x++
		return x
	}

}

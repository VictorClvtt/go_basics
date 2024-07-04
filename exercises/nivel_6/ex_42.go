package main

import "fmt"

func hello() func() {
	return func() {
		fmt.Println("Olha eu aqui!")
	}
}

func main() {

	x := hello()
	x()

}

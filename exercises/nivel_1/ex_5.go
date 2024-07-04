package main

import "fmt"

type meu_tipo int

var x meu_tipo
var y int

func main() {

	fmt.Printf("Valor de x: %v\n", x)
	fmt.Printf("Tipo de x: %T\n", x)

	x = 42

	fmt.Println("------------------------------")
	fmt.Printf("Novo valor de x: %v\n", x)
	fmt.Printf("Novo tipo de x: %T\n", x)
	fmt.Println("------------------------------")

	y = int(x)

	fmt.Printf("Valor de y: %v\n", y)
	fmt.Printf("Tipo de y: %T\n", y)
}

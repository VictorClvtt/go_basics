package main

import "fmt"

type meu_tipo int
var x meu_tipo

func main(){

	fmt.Printf("Valor: %v\n", x)
	fmt.Printf("Tipo: %T\n", x)

	x = 42

	fmt.Printf("Valor: %v\n", x)
	
}
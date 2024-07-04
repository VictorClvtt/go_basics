package main

import "fmt"

func main() {
	// Constante tipada
	const typedConst int = 42
	fmt.Println("Constante tipada:", typedConst)

	// Constante não tipada
	const untypedConst = "Hello, World!"
	fmt.Println("Constante não tipada:", untypedConst)
}

package main

import "fmt"

func main() {

	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["alfredo"])

	amigos["gopher"] = 4567764

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	// Declaração de variáveis feita junto com o teste lógico limitando as variaveis apenas ao escopo do if, fazendo com que elas deixem de existir ao final dos blocos.
	// sera recebe o valor do "indice" caso ele exista e ok recebe false caso o indice não exista
	// comma ok idiom
	if sera, ok := amigos["inexistente"]; !ok {
		fmt.Println("Não existe.")
	} else {
		fmt.Println(sera)
	}

}

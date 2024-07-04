package main

import "fmt"

func main() {

	// usando ponteiros, também é possível alterar o valor original de variáveis mesmo que elas sejam passadas como argumento em funções

	num := 0

	recebe_valor(num)
	fmt.Println("Na main:", num)

	// recebe o endereço de x
	recebe_ponteiro(&num)
	fmt.Println("Na main:", num)

}

func recebe_valor(x int) {
	x++
	fmt.Println("Na função por valor:", x)
}

func recebe_ponteiro(x *int) {
	// soma ao conteúdo do endereço do parâmetro
	*x++
	fmt.Println("Na função por ponteiro:", *x)
}

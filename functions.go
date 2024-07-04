package main

import "fmt"

// exemplo prático
// função variádica, o x é um slice(array) que recebe todos os parâmetros passados
func int_sum(x ...int) (int, int) {
	var sum int
	for _, v := range x {
		sum += v
	}

	return sum, len(x)
}

func main() {

	// exemplo de extrutura de uma função em go
	// func (receiver)                          identifier(parameters) (returns)           {code}
	//      anexa a função(método) a um tipo    nome       parâmetros  tipos dos retornos  código a ser executado

	// tudo em go é pass by value, ou seja as funções pegam uma copia das variável passadas como argumento mas não as variáveis em si
	num_1 := 20
	num_2 := 15

	fmt.Println(int_sum(num_1, num_2))
}

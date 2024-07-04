package main

import "fmt"

func funcao_argumento() {
	fmt.Println("Olha eu aqui de novo!")
}

func funcao_funcao(x func()) {
	fmt.Println("É O SEGUINTE:")
	x()
}

func main() {
	funcao_funcao(funcao_argumento)
}

package main

import "fmt"

func main() {

	x := retorna_uma_funcao()
	y := x(3)

	fmt.Println(y)

}

// função que retorna uma função que retorna um int
func retorna_uma_funcao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

package main

import "fmt"

func main() {

	// defer faz com que a linha na qual ele está escrito seja executada apenas no final do algorítmo
	defer fmt.Println("1")
	// quando mais de uma linha tem defer no algorítmo as linhas que aparecem por ultimo no algorítmo serão executadas primeiro
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")

}

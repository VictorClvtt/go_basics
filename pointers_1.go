package main

import "fmt"

func main() {

	// ponteiros são úteis para economizar memória
	// quando se trabalha com muitos valores, ao invés de copiar um valor com frequência para outras variáveis, é possível referenciar o valor dessa variável em outras, poupando espaço de memória

	x := 10

	// y recebe apenas o endereço de x, mas com isso é possivel acessar o valor de x colocando * antes de y
	y := &x

	fmt.Print("Valor de x: ")
	fmt.Println(x)
	fmt.Print("Endereço de x na memória: ")
	fmt.Println(&x)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Print("Valor de x a partir do ponteiro y: ")
	fmt.Println(*y)
	fmt.Print("Endereço de x na memória a partir do ponteiro y: ")
	fmt.Println(y)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("Tipo de x: %T\n", x)
	fmt.Printf("Tipo de y: %T\n", y)

}

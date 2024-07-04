package main

import "fmt"

func main() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "João",
		idade: 20,
	}

	fmt.Println(x.nome)
	fmt.Println(x.idade)

}

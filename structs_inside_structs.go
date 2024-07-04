package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	fumante   bool
}

type profissional struct {
	pessoa  pessoa
	titulo  string
	salario int
}

func main() {

	pessoa_1 := pessoa{
		nome:      "João",
		sobrenome: "Clivatti",
		idade:     20,
		fumante:   false,
	}

	pessoa_1_profissional := profissional{
		pessoa:  pessoa_1,
		titulo:  "Programador",
		salario: 100000,
	}

	pessoa_2 := pessoa{
		nome:      "Regina",
		sobrenome: "Rouca",
		idade:     60,
		fumante:   true,
	}

	fmt.Println(pessoa_1_profissional)
	fmt.Println(pessoa_2)

}

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	fumante   bool
}

func main() {
	pessoa_1 := pessoa{
		nome:      "João",
		sobrenome: "Clivatti",
		idade:     20,
		fumante:   false,
	}

	pessoa_2 := pessoa{
		nome:      "Regina",
		sobrenome: "Rouca",
		idade:     60,
		fumante:   true,
	}

	fmt.Println(pessoa_1)
	fmt.Println(pessoa_2)

}

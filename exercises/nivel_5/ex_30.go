package main

import "fmt"

type pessoa struct {
	nome                   string
	sobrenome              string
	sabores_fav_de_sorvete []string
}

func main() {
	pessoa_1 := pessoa{
		nome:                   "João",
		sobrenome:              "Clivatti",
		sabores_fav_de_sorvete: []string{"chocolate", "morango"},
	}

	pessoa_2 := pessoa{
		nome:                   "Carol",
		sobrenome:              "Filonzi",
		sabores_fav_de_sorvete: []string{"baunilha", "morango", "açaí"},
	}

	fmt.Printf("Nome: %v\n", pessoa_1.nome)
	fmt.Printf("Sobrenome: %v\n", pessoa_1.sobrenome)
	fmt.Printf("Sabores favoritos de sorvete: ")
	for i, v := range pessoa_1.sabores_fav_de_sorvete {

		if i == len(pessoa_1.sabores_fav_de_sorvete)-1 {
			fmt.Printf("%v\n", v)
		} else {
			fmt.Printf("%v, ", v)
		}

	}

	fmt.Println("---------------------------------------------------------")

	fmt.Printf("Nome: %v\n", pessoa_2.nome)
	fmt.Printf("Sobrenome: %v\n", pessoa_2.sobrenome)
	fmt.Printf("Sabores favoritos de sorvete: ")
	for i, v := range pessoa_2.sabores_fav_de_sorvete {

		if i == len(pessoa_2.sabores_fav_de_sorvete)-1 {
			fmt.Printf("%v\n", v)
		} else {
			fmt.Printf("%v, ", v)
		}

	}

}

package main

import "fmt"

type pessoa struct {
	nome                   string
	sobrenome              string
	sabores_fav_de_sorvete []string
}

func main() {

	mapa := make(map[string]pessoa)

	mapa["Pimentão"] = pessoa{
		nome:                   "João",
		sobrenome:              "Clivatti",
		sabores_fav_de_sorvete: []string{"chocolate", "morango"},
	}

	mapa["da Prússia"] = pessoa{
		nome:                   "Carol",
		sobrenome:              "Filonzi",
		sabores_fav_de_sorvete: []string{"baunilha", "morango", "açaí"},
	}

	for _, v := range mapa {
		fmt.Println(v)
	}

}

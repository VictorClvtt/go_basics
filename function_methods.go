package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

// anexa a função ao tipo pessoa
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {

	pessoa_1 := pessoa{
		nome:  "João",
		idade: 20,
	}

	pessoa_1.oibomdia()

}

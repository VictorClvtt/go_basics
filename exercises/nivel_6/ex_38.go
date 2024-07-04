package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) apresenta_pessoa() {
	fmt.Printf("Olá, meu nome é %v %v e eu tenho %v anos", p.nome, p.sobrenome, p.idade)
}

func main() {

	pessoa_1 := pessoa{
		nome:      "Victor",
		sobrenome: "Clivatti",
		idade:     20,
	}

	pessoa_1.apresenta_pessoa()

}

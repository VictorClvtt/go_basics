package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mude_me(p *pessoa) {
	p.nome = "Nada"
	p.sobrenome = "Nada"
	p.idade = 0
}

func main() {

	x := pessoa{
		nome:      "João",
		sobrenome: "Clivatti",
		idade:     20,
	}

	fmt.Println(x)
	mude_me(&x)
	fmt.Println(x)

}

package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	Carteira  float64
}

func main() {

	pessoa_1 := pessoa{
		Nome:      "Victor",
		Sobrenome: "Clivatti",
		Idade:     20,
		Profissao: "Programador",
		Carteira:  9500000,
	}

	pessoa_2 := pessoa{"Carol", "Filonzi", 18, "Psicóloga", 1000000000}

	p_1, err := json.Marshal(pessoa_1)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println(string(p_1))

	p_2, err := json.Marshal(pessoa_2)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println(string(p_2))

}

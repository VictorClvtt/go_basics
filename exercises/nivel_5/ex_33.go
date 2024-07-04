package main

import "fmt"

func main() {

	x := struct {
		filmes_assistidos []string
		amigos_e_idades   map[string]int
	}{
		filmes_assistidos: []string{
			"O Poderoso Chefão",
			"Brilho Eterno de Uma Mente sem Lembranças",
		},
		amigos_e_idades: map[string]int{
			"João":      20,
			"Carol":     18,
			"Suzy":      5,
			"Pepeto":    1,
			"Gatão":     15,
			"Retalinho": 13,
			"Skye":      7,
		},
	}

	fmt.Println(x)

}

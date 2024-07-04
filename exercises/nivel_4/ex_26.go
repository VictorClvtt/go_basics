package main

import "fmt"

func main() {

	x := [][]string{
		[]string{"Victor", "Clivatti", "Estudar"},
		[]string{"Carol", "Filonzi", "Nhenhenhe e Estudar"},
		[]string{"Gatão", "Clivatti", "Dormir"},
	}

	for _, v := range x {
		fmt.Printf("Nome: %v | Sobrenome: %v | Hobby: %v\n", v[0], v[1], v[2])
	}

}

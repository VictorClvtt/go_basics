package main

import (
	"fmt"
)

func main() {

	diasDaSemana := [7]string{"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"}
	hoje := diasDaSemana[2]

	switch hoje {
	case "Domingo":
		fmt.Println("Hoje é: Domingo")
	case "Segunda-feira":
		fmt.Println("Hoje é: Segunda-feira")
	case "Terça-feira":
		fmt.Println("Hoje é: Terça-feira")
	case "Quarta-feira":
		fmt.Println("Hoje é: Quarta-feira")
	case "Quinta-feira":
		fmt.Println("Hoje é: Quinta-feira")
	case "Sexta-feira":
		fmt.Println("Hoje é: Sexta-feira")
	case "Sábado":
		fmt.Println("Hoje é: Sábado")
	default:
		fmt.Println("Dia da semana inválido")
	}
}

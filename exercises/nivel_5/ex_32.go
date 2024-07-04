package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo             veiculo
	tracao_quatro_rodas bool
}

type sedan struct {
	veiculo     veiculo
	modelo_luxo bool
}

func main() {

	caminhonete_1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Azul",
		},
		tracao_quatro_rodas: true,
	}

	sedan_1 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Vermelho",
		},
		modelo_luxo: true,
	}

	fmt.Println(caminhonete_1)
	fmt.Println(sedan_1)

}

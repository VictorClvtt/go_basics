package main

import "fmt"

func main() {

	combustivel := 76

	if combustivel <= 25 {
		fmt.Println("Abastecer urgentemente.")
	} else if combustivel > 25 && combustivel <= 75 {
		fmt.Println("Abastecer não tão urgentemente.")
	} else {
		fmt.Println("Quantidade satisfatória de combustível.")
	}

}

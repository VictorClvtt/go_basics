package main

import "fmt"

func main() {

	slice := []string{"banana", "maçã", "jaca"}

	for index, value := range slice {
		fmt.Printf("No índice %v temos o valor %v\n", index, value)
	}

	slice[3] = "melancia"

	for index, value := range slice {
		fmt.Printf("No índice %v temos o valor %v\n", index, value)
	}

}

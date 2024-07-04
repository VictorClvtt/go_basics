package main

import "fmt"

func main() {

	slice := []string{"pepperoni", "mussarela", "quatro queijos", "brigadeiro"}

	slice_slice := slice[2:len(slice)]

	fmt.Println(slice_slice)

}

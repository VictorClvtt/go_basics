package main

import "fmt"

func main() {

	slice_1 := []string{"pepperoni", "mussarela", "quatro queijos"}
	fmt.Println(slice_1)

	slice_1 = append(slice_1, "brigadeiro")
	fmt.Println(slice_1)

	slice_2 := []string{"marguerita", "M&M's"}

	slice_1 = append(slice_1, slice_2...)
	fmt.Println(slice_1)

}

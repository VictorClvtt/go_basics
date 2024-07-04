package main

import "fmt"

func main() {

	amigos := map[string]int{
		"joão":   5551234,
		"carol":  9996674,
		"gopher": 9873426,
	}

	for k, v := range amigos {
		fmt.Println(k, v)
	}

	delete(amigos, "gopher")
	fmt.Println("---------------")

	for k, v := range amigos {
		fmt.Println(k, v)
	}

	amigos["suzy"] = 6298267
	fmt.Println("---------------")

	for k, v := range amigos {
		fmt.Println(k, v)
	}

}

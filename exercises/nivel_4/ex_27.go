package main

import "fmt"

func main() {

	x := map[string]string{
		"Victor_Clivatti": "dormir",
		"Carol_Filonzi":   "nhenhenhe",
	}

	for k, v := range x {
		fmt.Printf("Nome: %v | Hobby: %v\n", k, v)
	}

}

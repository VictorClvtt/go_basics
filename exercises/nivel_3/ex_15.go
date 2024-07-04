package main

import "fmt"

func main() {

	nascimento := 2003
	ano_meta := 2077

	for {
		fmt.Println(nascimento)
		nascimento++

		if nascimento > ano_meta {
			break
		}
	}
}

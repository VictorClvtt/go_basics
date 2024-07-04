package main

import "fmt"

func main() {

	nascimento := 2003
	ano_meta := 2077

	for nascimento <= ano_meta {
		fmt.Println(nascimento)
		nascimento++
	}
}

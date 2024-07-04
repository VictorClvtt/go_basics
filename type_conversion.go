package main

import "fmt"

type hotdog int
var b hotdog = 100

func main(){

	x := 10
	fmt.Printf("%v\n", x)

	x = int(b)
	fmt.Printf("%v\n", x)

}
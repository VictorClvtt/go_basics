package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main(){

	s := fmt.Sprintf("x(valor: %v, tipo: %T)\n", x, x)
	s += fmt.Sprintf("y(valor: %v, tipo: %T)\n", y, y)
	s += fmt.Sprintf("z(valor: %v, tipo: %T)\n", z, z)
	
	fmt.Print(s)

}
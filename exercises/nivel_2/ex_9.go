package main

import "fmt"

func main() {

	x := 123
	fmt.Printf("x(Decimal: %v, Binário: %b, Hexadecimal: %X)\n", x, x, x)

	y := x << 1
	fmt.Printf("y(Decimal: %v, Binário: %b, Hexadecimal: %X)\n", y, y, y)

}

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.altura * q.largura
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) {
	fmt.Println(f.area())
}

func main() {

	quadrado_1 := quadrado{
		altura:  10,
		largura: 25,
	}

	info(quadrado_1)

}

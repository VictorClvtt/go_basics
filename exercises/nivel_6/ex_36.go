package main

import "fmt"

func main() {

	x := []int{2, 5, 8, 2, 4, 9}

	fmt.Println(sum(x...))
	fmt.Println(sum_slice(x))

}

func sum(x ...int) int {
	var sum int
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func sum_slice(x []int) int {
	var sum int
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

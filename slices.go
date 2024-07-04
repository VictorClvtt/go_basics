package main

import "fmt"

func main() {

	array := [5]int{1, 4, 6, 3, 9}
	fmt.Println(array)

	slice := []int{10, 4, 3, 4, 9, 13}
	fmt.Println(slice)
	slice = append(slice, 7)
	slice[3] = 123
	fmt.Println(slice)

}

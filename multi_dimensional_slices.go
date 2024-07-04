package main

import "fmt"

func main() {

	slice_1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(slice_1)

	fmt.Println("-------------------------")

	for i := 0; i < len(slice_1); i++ {
		fmt.Println(slice_1[i])
	}

	fmt.Println("-------------------------")

	for i := 0; i < len(slice_1); i++ {
		for j := 0; j < len(slice_1[i]); j++ {
			fmt.Println(slice_1[i][j])
		}
	}

}

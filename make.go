package main

import "fmt"

func main() {

	slice_1 := make([]int, 5, 10)

	slice_1[0], slice_1[1], slice_1[2], slice_1[3], slice_1[4] = 1, 2, 3, 4, 5

	slice_1 = append(slice_1, 6)

	fmt.Println(slice_1, len(slice_1), cap(slice_1))

	slice_1 = append(slice_1, 7)
	slice_1 = append(slice_1, 8)
	slice_1 = append(slice_1, 9)
	slice_1 = append(slice_1, 10)

	fmt.Println(slice_1, len(slice_1), cap(slice_1))

	slice_1 = append(slice_1, 10)

	fmt.Println(slice_1, len(slice_1), cap(slice_1))

}

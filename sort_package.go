package main

import (
	"fmt"
	"sort"
)

func main() {
	string_slice := []string{"João", "Carol", "Gatão", "Retalho", "Suzy", "Pepeto", "Skye"}
	int_slice := []int{24, 63, 15, 11, 29, 43, 46, 58, 33, 60, 59}

	fmt.Println(string_slice)
	sort.Strings(string_slice)
	fmt.Println(string_slice)

	fmt.Println("---------------------------------------------")

	fmt.Println(int_slice)
	sort.Ints(int_slice)
	fmt.Println(int_slice)
}

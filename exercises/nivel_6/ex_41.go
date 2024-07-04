package main

import "fmt"

func hello(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}

func main() {
	x := hello("João")
	fmt.Println(x)
}

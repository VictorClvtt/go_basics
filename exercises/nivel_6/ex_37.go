package main

import "fmt"

func main() {

	defer fmt.Println("Apareço primeiro no código mas sou executado depois.")
	fmt.Println("Apareço depois no código mas sou executado primeiro.")

}

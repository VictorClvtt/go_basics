package main

import "fmt"

func main(){

	// Go short declaration operator method
	// This method automatically assigns the type based on the type of the value assigned when initialized
	variavel1 := 123
	fmt.Println(variavel1)

	// Go long declaration method
	// In this method the type of the variable is hard coded making it impossible to eventually recieve an unexpected data type and generating errors
	var variavel2 int = 321
	fmt.Println(variavel2)

}
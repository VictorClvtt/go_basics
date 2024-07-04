package main

import "fmt"

const (
    a = 1994 + iota
    b
    c
    d
    e
)

func main() {
	
    fmt.Printf("%v, %v, %v, %v", b, c, d, e)

}
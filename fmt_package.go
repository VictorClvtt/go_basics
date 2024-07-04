package main

import "fmt"

func main() {
    // Print examples
    fmt.Println("Hello, Go!")  // Print with newline
    fmt.Print("Hello, ")       // Print without newline
    fmt.Println("Go!")         // Print with newline
    
    // Print formatted output
    fmt.Printf("Hello, %s!\n", "world") // Formatted print with newline
    fmt.Printf("Value: %d\n", 42)       // Formatted print integer
    
    // Print types and variables
    fmt.Printf("Type of 42 is %T\n", 42)      // Print type
    fmt.Printf("Value of pi is %.2f\n", 3.14159265359) // Print float with precision
    
    // Print boolean values
    fmt.Printf("It is %t that Go is awesome!\n", true) // Print boolean
    
    // Print pointers
    a := 10
    fmt.Printf("Address of a: %p\n", &a) // Print pointer address
    
    // Print multiple values
    fmt.Println("Values:", 1, 2, 3) // Print multiple values
    
    // Print newline explicitly
    fmt.Print("This ")
    fmt.Print("is ")
    fmt.Println("on ")
    fmt.Println("multiple lines.") // Newline control
    
    // Sprint, Sprintf, Sprintln examples
    s := fmt.Sprint("Hello, ", "Go!")
    fmt.Println("Sprint:", s)
    
    s2 := fmt.Sprintf("Value: %d", 42)
    fmt.Println("Sprintf:", s2)
    
    s3 := fmt.Sprintln("Hello,", "Go!")
    fmt.Println("Sprintln:", s3)
}
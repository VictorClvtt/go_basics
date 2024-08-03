package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define the struct that matches the JSON structure
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {
	// Example JSON data
	jsonData := `{
        "name": "John Doe",
        "age": 30,
        "email": "john.doe@example.com",
        "address": "123 Main St"
    }`

	// Create an instance of the struct
	var person Person

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print the struct to verify the data
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Address: %s\n", person.Address)
}

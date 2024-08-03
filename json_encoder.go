package main

import (
	"encoding/json"
	"os"
)

func main() {
	// Define a anonymous struct to encode
	jamesbond := struct {
		Name  string
		Agent int
	}{
		Name:  "James Bond",
		Agent: 007,
	}

	// Create a new encoder that writes to standard output
	encoder := json.NewEncoder(os.Stdout)

	// Encode the struct as JSON and write it to standard output
	encoder.Encode(jamesbond)
}

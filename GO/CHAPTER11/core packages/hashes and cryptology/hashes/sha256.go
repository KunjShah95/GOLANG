package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// Input data
	data := "Hello, Go!"

	// Create a new SHA256 hash
	hash := sha256.New()

	// Write data to the hash
	hash.Write([]byte(data))

	// Calculate the final hash
	hashBytes := hash.Sum(nil)

	// Convert hash to a hex string
	hashString := hex.EncodeToString(hashBytes)

	// Output the hash
	fmt.Println("SHA256 Hash:", hashString)
}

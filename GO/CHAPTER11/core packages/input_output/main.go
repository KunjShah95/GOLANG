package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Step 1: Input - Read user input
	var name string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	// Step 2: Output - Write to a file
	filename := "greeting.txt"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write a greeting message to the file
	_, err = fmt.Fprintf(file, "Hello, %s!\n", name)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Greeting written to", filename)

	// Step 3: Input - Read from the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Step 4: Output - Print the content of the file
	fmt.Println("Content of the file:")
	fmt.Println(string(data))
}

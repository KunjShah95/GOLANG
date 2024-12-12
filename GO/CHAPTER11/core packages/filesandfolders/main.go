package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Step 1: Create a directory
	dirName := "example_dir"
	err := os.Mkdir(dirName, 0755) // 0755 is the permission mode
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Directory already exists:", dirName)
		} else {
			log.Fatal("Error creating directory:", err)
		}
	} else {
		fmt.Println("Directory created:", dirName)
	}

	// Step 2: Create a file in the directory
	fileName := dirName + "/greeting.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write a greeting message to the file
	_, err = fmt.Fprintf(file, "Hello, World!\n")
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
	fmt.Println("File created and written to:", fileName)

	// Step 3: Read the content of the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Step 4: Print the content of the file
	fmt.Println("Content of the file:")
	fmt.Println(string(data))

	// Step 5: List files in the directory
	fmt.Println("Files in the directory:")
	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

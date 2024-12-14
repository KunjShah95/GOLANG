package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Create a new file
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Write to the file
	d := []byte("Hello, World!")
	f.Write(d)

	// Read from the file
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// Check if the file exists
	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}

	// Delete the file
	err = os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File deleted")
}

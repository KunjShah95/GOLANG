package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}
	person := Person{"Max", 20}
	fmt.Println(person)
}

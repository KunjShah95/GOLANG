package main

import "fmt"

func main() {
	person := struct {
		name string
		age  int
	}{name: "John", age: 30}
	fmt.Println(person.name) // prints "John"
	fmt.Println(person.age)  // prints 30
}

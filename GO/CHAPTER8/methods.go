// A method is a function that is declared on a type.
// It is used to add behavior to a type.

// Example
package main

import "fmt"

type Person struct {
	name string
	age  int
}

// A method on the Person type
func (p Person) speak() {
	fmt.Printf("My name is %s, and I am %d years old.\n", p.name, p.age)
}

func main() {
	p := Person{"John", 30}
	p.speak() // prints "My name is John, and I am 30 years old."
}

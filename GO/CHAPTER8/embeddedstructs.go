package main

import "fmt"

func main() {
	type Address struct {
		street string
		city   string
	}

	type Person struct {
		name string
		age  int
		Address
	}

	person := Person{name: "John", age: 30, Address: Address{street: "123 Main St", city: "Anytown"}}
	fmt.Println(person.name)   // prints "John"
	fmt.Println(person.age)    // prints 30
	fmt.Println(person.street) // prints "123 Main St"
	fmt.Println(person.city)   // prints "Anytown"
}

package main

import "fmt"

// Define a struct called Person
type Person struct {
	// Field 1: name of type string
	name string
	// Field 2: age of type int
	age int
	// Field 3: address of type Address
	address Address
}

// Define a struct called Address
type Address struct {
	// Field 1: street of type string
	street string
	// Field 2: city of type string
	city string
}

func main() {
	// Create a new Person instance
	person := Person{
		// Initialize the name field
		name: "John",
		// Initialize the age field
		age: 30,
		// Initialize the address field
		address: Address{
			// Initialize the street field
			street: "123 Main St",
			// Initialize the city field
			city: "Anytown",
		},
	}

	// Access the fields of the person instance
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("Street:", person.address.street)
	fmt.Println("City:", person.address.city)
}

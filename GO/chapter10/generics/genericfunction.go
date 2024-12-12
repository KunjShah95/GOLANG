package main

import "fmt"

// generic function in go
func identity[T any](arg T) T {
	return arg
}

func main() {
	// generic function in go
	fmt.Println(identity[int](5))         // prints 5
	fmt.Println(identity[string]("five")) // prints "five"
}

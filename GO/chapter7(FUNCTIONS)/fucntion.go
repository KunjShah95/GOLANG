package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	//wrong method to declare function as you can't declare a function in a function itself
	// func greetertwo(){
	// 	fmt.Println("another method")
	// }

	result := add(3, 5)

	fmt.Println("Result is:", result)

	// Call the proadd function with a slice of values
	values := []int{1, 2, 3, 4, 5}
	result = proadd(values...)
	fmt.Println("Result of proadd is:", result)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proadd(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}

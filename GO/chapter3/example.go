package main

import "fmt"

func main() {
	fmt.Print("Enter a number:  ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Printf("The double of %f is %f\n", input, output)
}

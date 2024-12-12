package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in golang")

	fruitList := []string{"Apple", "Mango", "Peach", "Watermelon"}
	fmt.Printf("Fruit List: %v\n", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Length of fruitList is", len(fruitList))
	fmt.Println("Capacity of fruitList is", cap(fruitList))
}

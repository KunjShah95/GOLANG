package main

import "fmt"

func findIndex(numbers []int, target int) int {
	for i, num := range numbers {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(findIndex(numbers, 3)) // prints 2
	fmt.Println(findIndex(numbers, 6)) // prints -1
}

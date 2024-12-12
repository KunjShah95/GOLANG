/* In Go, a variadic function is a function that can accept a variable number of arguments. This is achieved by using the ... syntax in the function signature. */

package main

import "fmt"

func sum(numebrs ...int) int {
	sum := 0

	for _, num := range numebrs {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // prints 15

}

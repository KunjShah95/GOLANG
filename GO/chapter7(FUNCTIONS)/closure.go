/* A closure is a function that:
Has access to its own scope.
Has access to the scope of its outer functions.
Can be returned from a function and used later.
Example of a Closure */

package main

import "fmt"

func outer() func() int {
	x := 10
	return func() int {
		x++
		return x
	}
}

func main() {
	inner := outer()
	fmt.Println(inner())
	fmt.Println(inner())
	fmt.Println(inner())
}

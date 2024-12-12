// recover is a built-in function in Go that can be used to regain control of a panicking program. It's commonly used in conjunction with defer to catch and handle panics.
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Start")
	panic("Something went wrong")
	fmt.Println("End")
}

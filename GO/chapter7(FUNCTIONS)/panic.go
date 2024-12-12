//panic is a built-in function in Go that stops the normal execution of a program and begins panicking. When a panic occurs, the program will stop executing and start unwinding the stack, calling any deferred functions it encounters.

package main

import "fmt"

func main() {
	fmt.Println("Start")
	panic("Something went wrong")
	fmt.Println("End")
}

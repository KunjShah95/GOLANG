//defer is a keyword in Go that schedules a function call to be executed when the surrounding function returns. It's commonly used to release resources, such as closing files or network connections, after they are no longer needed.

package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Deferred function")
	fmt.Println("End")
}

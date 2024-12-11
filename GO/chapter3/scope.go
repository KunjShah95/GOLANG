// package main

// import "fmt"

// func main() {
// 	var x string = "Hello World"
// 	fmt.Println(x)
// }

//2nd method
// package main

// import "fmt"

// var x string = "Hello World"

// func main() {
// 	fmt.Println(x)
// }

/*Notice that we moved the variable outside of the main
function. This means that other functions can access
this variable */

// package main

// import "fmt"

// var x string = "Hello World"

// func main() {
// 	fmt.Println(x)
// }
// func f() {
// 	fmt.Println(x)
// }

//another emthod to do above thing as same
package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
}
func f() {
	fmt.Println(x)
}

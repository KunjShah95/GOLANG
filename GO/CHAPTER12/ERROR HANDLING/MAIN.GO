package main

import "fmt"

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// another method to handle error with strings

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "cannot divide by zero"
	}
	return a / b, "nil"
}
func main() {
	fmt.Println("error handling!")
	ans, _ := divide(10, 2)
	//another method to handle error
	// ans, err := divide(10, 4)
	// fmt.Println(err)
	fmt.Println(ans)
}

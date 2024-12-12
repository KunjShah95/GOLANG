package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}

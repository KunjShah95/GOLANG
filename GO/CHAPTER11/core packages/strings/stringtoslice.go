package main

import "strings"

func main() {
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(arr) // [116 101 115 116]
	fmt.Println(str) // test

}

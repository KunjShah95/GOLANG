package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("learning web services")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET respond", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Type of response :%T\n", res)
	// fmt.Println("response: ", res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("eror reading response: ", err)
		return
	}
	fmt.Println("response: ", string(data))
}

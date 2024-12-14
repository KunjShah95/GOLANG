package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response: ", res.Status)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Todo: ", todo)
	fmt.Println("Title Response: ", todo.Title)
	fmt.Println("Completed: ", todo.Completed)
}

func performPostRequest() {
	var todo = Todo{
		UserID:    23,
		Title:     "Kunj Shah",
		Completed: true,
	}
	//convert the todo struct into json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		return
	}

	//convert the json data into string
	jsonString := string(jsonData)

	//convert the string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	//send POST request
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response :", err)
		return
	}
	fmt.Println("Response: ", string(data))
}

func performUpdateRequest() {
	var todo = Todo{
		UserID:    23,
		Title:     "Kunj Shah",
		Completed: false,
	}
	//convert the todo struct into json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		return
	}

	//convert the json data into string
	jsonString := string(jsonData)

	//convert the string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("Error in creating put request :", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	//send PUT request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response :", err)
		return
	}
	fmt.Println("Response: ", string(data))
	fmt.Println("Response Status: ", res.Status)
}

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	//create DELETE request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error in creating delete request :", err)
		return
	}
	//send the http request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response Status: ", res.Status)
}

func main() {
	fmt.Println("Learning Crud")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}

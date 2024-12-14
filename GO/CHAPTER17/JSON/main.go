package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("learning json")

	person := Person{"John", 30, true}
	fmt.Println(person)

	//convert person into json encoding (marshalling)

	jsonPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}
	fmt.Println(string(jsonPerson))

	//decoding(unmarshalling)

	var p Person

	err = json.Unmarshal(jsonPerson, &p)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
		return
	}

	fmt.Println(p)
}

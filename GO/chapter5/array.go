package main

import "fmt"

func main() {
	fmt.Println("Welcoem to array in golang.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"
	fmt.Println("Fruit List:", fruitList)

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg List:", vegList)

	fmt.Println("Fruiut list is:", len(fruitList))

	fmt.Println("Vegetable list is:", len(vegList))
}

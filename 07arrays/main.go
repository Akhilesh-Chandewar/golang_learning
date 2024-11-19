package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [6]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"
	fruitList[5] = "Orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", fruitList[1])
	fmt.Println("Fruit list is: ", fruitList[4])
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegetableList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegetable list is: ", vegetableList)
	fmt.Println("Vegetable list is: ", len(vegetableList))
}


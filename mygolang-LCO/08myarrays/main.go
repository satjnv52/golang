package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	var fruitList = [5]string{"Apple", "Mango", "Potato", "Tomato", "Berries"}

	fmt.Println("Print element of fruitlist: ", fruitList)
	fmt.Println("Length of array is :", len(fruitList))
}

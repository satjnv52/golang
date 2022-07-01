package main

import "fmt"

func main() {
	fmt.Println("Welcome to concept of pointers")
	//Whenever we declare some var and try to use it, copy of this var is created from memory
	//and shared. Sometimes in case of complexity, there might be some irrelugarities in data sharing
	//pointers provides us assurity of passing value. It share data using memory addrs. All operations
	//perform actuallly on values not their copies
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 34
	var ptr = &myNumber
	fmt.Println("value of pointer is ", ptr)
	fmt.Println("value of pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value after operation thr pointer:", myNumber)

}

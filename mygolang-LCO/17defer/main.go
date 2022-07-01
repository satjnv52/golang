package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

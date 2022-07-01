package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in go lang")
	count := 23
	token := 0
	tokenOdd := 0
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			token++
		} else {
			tokenOdd++
		}
	}
	fmt.Println("Number of even number is:", token)
	fmt.Println("Number of odd number:", tokenOdd)

}

package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeting()
	sum := adder(3, 5)
	result, msg := proAdder(2, 3, 4, 5)
	fmt.Println("Result of pro function is, ", result)
	fmt.Println("Message from ProAdder is:", msg)
	fmt.Println("summation is", sum)
}

func greeting() {
	fmt.Println("Namasate from India")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for i := range values {
		total += values[i]
	}
	str := "Hi"
	return total, str
}

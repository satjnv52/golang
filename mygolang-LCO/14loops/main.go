package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thusday"}
	fmt.Println(days)

	for i := range days {
		if days[i] == "Sunday" {
			fmt.Println("Its:", days[i])
			//break
		} else if days[i] == "Monday" {
			goto day
		} else {
			fmt.Println("Its not sunday")
		}
	}

day:
	fmt.Println("Its Mondays using goto")
}

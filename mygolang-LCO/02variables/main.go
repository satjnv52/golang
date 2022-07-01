package main

import "fmt"

const jswToken = "LoggedInSuccessfully"

func main() {
	var username string = "SATYAM"
	fmt.Println(username)
	fmt.Printf("Type of variable is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable is %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Printf("this is unsigned integer %d \n", smallValue)
	fmt.Printf("Type of variable is %T \n", smallValue)

	var anotherVar bool
	fmt.Println(anotherVar)
	fmt.Printf("Type of variable is %T \n", anotherVar)

	fmt.Println(jswToken)
	fmt.Printf("Type of variable is %T \n", jswToken)

}

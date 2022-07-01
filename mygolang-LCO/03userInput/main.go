package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user Feedback"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for Blue Ridge:")

	//comma ok syntex
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating! \nRating of Blue Ridge is ", input)
	fmt.Printf("Type of this rating is %T", input)
}

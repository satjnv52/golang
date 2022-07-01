package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in Golang")

	progLang := make(map[string]string)
	progLang["JS"] = "Javascript"
	progLang["RB"] = "Ruby"
	progLang["PY"] = "Python"

	fmt.Println("list of languages are:", progLang)

	delete(progLang, "RB")
	fmt.Println("list of languages after delete:", progLang)

	//for loop
	for i, j := range progLang {
		fmt.Printf("For key %v value is %v \n", i, j)
	}

}

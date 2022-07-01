package main

import "fmt"

func main() {
	fmt.Println("Wlecome to structs in golang")

	sat := User{"Satyam", "sat@google.com", 23, true}
	fmt.Println(sat)
	fmt.Printf("Details of satyam is %+v \n", sat)
	fmt.Printf("Name of first user is %v and his email is %v.\n", sat.Name, sat.Email)

	sat.getName()
	sat.getEmail()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getName() {
	fmt.Println("Name of User is:", u.Name)
}

func (u User) getEmail() {
	u.Email = "newEmail@go.dev"
	fmt.Println("New Enail is:", u.Email)
}

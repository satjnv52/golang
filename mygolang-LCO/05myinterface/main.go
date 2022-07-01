package main

import "fmt"

//creating interface
type myInterface interface {
	//methods
	Tarea() float64
	volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

//implementing methods of tank interface

func (m myvalue) Tarea() float64 {
	return 2*3.14*m.radius*m.radius + 2*m.radius*m.height
}

func (m myvalue) volume() float64 {
	return (3.14 * m.radius * m.radius * m.height)
}
func main() {
	fmt.Println("Welcome to concept of interface in Golang")
	var t myInterface // we can create variable of interface, we cant intanticate intrfc in golang
	t = myvalue{10, 14}
	fmt.Println("Total Area is:", t.Tarea())
	fmt.Println("Total Voume is:", t.volume())
	fmt.Printf("Type of myInterfqace is: %T", t)

}

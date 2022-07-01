package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println("Switches in Go")
	rand.Seed(time.Now().Unix())
	outCome := rand.Intn(7)
	switch outCome {
	case 0:
		fmt.Println("Number is 0")
		fallthrough
	case 1:
		fmt.Println("Number is 1")
		break

	case 2:
		fmt.Println("Number is 2")
		break
	default:
		{
			fmt.Println("Case is Default")
		}
	case 3:
		{
			fmt.Println("Number is 3")
		}
	case 4:
		{
			fmt.Println("Number is 4")
		}
	case 5:
		{
			fmt.Println("Number is 5")
		}
	case 6:
		{
			fmt.Println("Number is 6")
		}
	}
}

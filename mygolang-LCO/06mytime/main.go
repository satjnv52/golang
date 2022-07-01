package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time Study in Go")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createDate := time.Date(2020, time.April, 24, 12, 12, 34, 2, time.Now().Location())
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

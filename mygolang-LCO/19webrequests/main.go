package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handeling web request using Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response is:", response)
	fmt.Printf("Response is of type: %T.\n", response)
	defer response.Body.Close() // callers respo to close the connections

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data type of reading is: %T. \n", databytes)
	fmt.Println("Content of reading is:", string(databytes))
}

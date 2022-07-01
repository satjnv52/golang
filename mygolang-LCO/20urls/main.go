package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=As342gBsdZy"

func main() {
	fmt.Println("Handelling URLS using Golang")
	fmt.Println("URL is:", myurl)

	//parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Host Name is:", result.Hostname())
	// fmt.Println("Scheme Name is:", result.Scheme)
	// fmt.Println("Host is:", result.Host)
	// fmt.Println("Path is:", result.Path)
	// fmt.Println("Port is:", result.Port())
	fmt.Println("Query is:", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of Query is: %T \n", qparams)
	fmt.Println(qparams["coursename"], qparams["paymentid"])

	for i, val := range qparams {
		fmt.Printf("Showing Key and Values of params: %v : %v \n", i, val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutorials",
		RawPath: "user-Satyam",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println("Another URL is:", anotherUrl)
}

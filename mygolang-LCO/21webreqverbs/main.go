package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Verb Tutorial")
	//PerfromGetRequest()
	//PerfromPostJsonRequest()
	PerfromPostFormRequest()
}

func PerfromGetRequest() {
	const myural = "http://localhost:8000/get"

	response, err := http.Get(myural)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // to close body of response//
	fmt.Printf("Type of response is: %T\n", response)
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var newStrBuild strings.Builder // builder is mth of creating string using write method//
	res, err := ioutil.ReadAll(response.Body)

	byteData, _ := newStrBuild.Write(res)

	if err != nil {
		panic(err)
	}
	fmt.Println("ByteData is: ", byteData)
	fmt.Println(newStrBuild.String())
	//fmt.Println("Resonse thr ReadAll:", string(res))
}

func PerfromPostJsonRequest() {
	const myural = "http://localhost:8000/post"

	//creating psuedo json payload

	requestBody := strings.NewReader(`
	{
		"Name" : "Satyam",
		"Coursename" : "Go with Golang",
		"Price" : "Free"
	}
	`)

	response, err := http.Post(myural, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content is : ", string(content))
}

func PerfromPostFormRequest() {
	const myurl = "http://localhost:8000/post"

	//create pseudo form data

	data := url.Values{}
	data.Add("fName", "Satyam")
	data.Add("lName", "Dubey")
	data.Add("eMail", "sat@google.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

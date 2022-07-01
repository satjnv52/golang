package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file system in golang")
	content := "Content of my first file which is cretaed by me"

	file, err := os.Create("./myfiles.txt")

	CheckNilError(err)

	length, err := io.WriteString(file, content)
	CheckNilError(err)

	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./myfiles.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)
	//reading of file is databyte not in string
	//for creation of file we need OS, for writing something we need inpt outpt io.write with file name
	//along with address
	// for reading also we need file name with address but in ioutil pacakage.
	CheckNilError(err)
	fmt.Println("Content in bytes in files:", databyte)
	fmt.Println("Content using string keyword:", string(databyte))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

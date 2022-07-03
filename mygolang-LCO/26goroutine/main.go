package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	//fmt.Println("Namaste! Wlecome for Goroutine")
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

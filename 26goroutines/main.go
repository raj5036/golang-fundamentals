package main

import (
	"fmt"
	"net/http"
	"sync"
)

// https://www.golang-book.com/books/intro/10
var wg sync.WaitGroup // This is actually a Pointer
var mut sync.Mutex    // Pointer
var signals = []string{}

func main() {
	fmt.Println("Handling goroutines")

	websites := make([]string, 3)
	websites[0] = "https://google.com"
	websites[1] = "https://lco.dev"
	websites[2] = "https://go.dev"

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}
	// go greeter("Hello")
	// greeter("World")
	// var s string
	// fmt.Scanln(&s)
	wg.Wait()
	fmt.Println("Signals", signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There is an error while fetching the endpoint")
		return
	}
	mut.Lock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	signals = append(signals, endpoint)
	mut.Unlock()
}

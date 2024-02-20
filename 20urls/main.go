package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=234j23gjd"

func main() {
	fmt.Println("Handling URLs here!")

	// Parse URL
	data, _ := url.Parse(myUrl)
	fmt.Println(data.Scheme)
	fmt.Println(data.Host)
	fmt.Println(data.Path)
	fmt.Println(data.Port())
	fmt.Println(data.RawQuery)

	queryParams := data.Query()
	fmt.Println("queryParams", queryParams, queryParams["coursename"])
}

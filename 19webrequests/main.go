package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Type of response %T\n", response)
		defer response.Body.Close() // This is coder's responsibility to close this API call
		dataBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response", string(dataBytes))
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web requests")
	// performGetRequest()
	// performPostJSONRequest()
	performPostFormRequest()
}

func performGetRequest() {
	const URL string = "http://localhost:3001"

	response, err := http.Get(URL)
	defer response.Body.Close() // VVI

	if err != nil {
		panic(err)
	}
	fmt.Println("StatusCode", response.StatusCode)
	fmt.Println("Response content length", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is", string(content))
}

func performPostJSONRequest() {
	const URL string = "http://localhost:3001/post"

	// fake JSON Payload
	requestBody := strings.NewReader(`
		{
			"username": "Raj GO",
			"password": "P@ssw0rd"
		}
	`)

	response, _ := http.Post(URL, "application/json", requestBody)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is", string(content))
}

func performPostFormRequest() {
	const URL string = "http://localhost:3001/post-form"

	// formData
	data := url.Values{}
	data.Add("firstname", "Raj")
	data.Add("age", "30")

	response, _ := http.PostForm(URL, data)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is", string(content))
}

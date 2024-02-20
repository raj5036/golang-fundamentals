// https://youtu.be/SZ5xZ9OTeEI?si=WZU6oer9oQfeMds0
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJSON() {
	lcoCourses := []course{
		{"React", 28, "Twitter", "password", []string{"1", "2"}},
		{"ME js", 285, "Youtube", "password123", []string{"1", "2"}},
		{"Angular", 228, "Twitter", "password123", nil},
	}

	// Package this data as JSON
	finalJSON, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React",
			"Price": 28,
			"Platform": "Twitter",
			"tags": ["1","2"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was Invalid")
	}

	// Some cases where you only want to put key value pair
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key: %v -> Value: %v & Type is %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Handling JSON response received from web request")
	// EncodeJSON()
	DecodeJSON()
}

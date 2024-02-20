package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")
	var fruits = [4]string{"Apple", "Banana", "Orange", "Guava"}
	fruits[3] = "Hmm"

	fmt.Println(fruits)
	fmt.Println(len(fruits))
}

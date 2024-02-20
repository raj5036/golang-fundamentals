package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// Structs doesn't offer Super or Inheritance related concepts
	Raj := User{"Raj Karmakar", "raj@go.dev", true}
	fmt.Println("Raj", Raj)
	fmt.Printf("Raj details are %+v\n", Raj)
	fmt.Printf("Raj name is %v and email is %v\n", Raj.Name, Raj.Email)
}

type User struct {
	Name       string
	Email      string
	IsVerified bool
}

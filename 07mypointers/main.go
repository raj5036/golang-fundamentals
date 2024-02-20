// https://youtu.be/-BFJ0dZyxHg?si=K1jLmVDMeQNRQGyp
package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers")

	var ptr *int
	fmt.Println("Value of ptr is", ptr)

	myNumber := 23
	var ptr1 = &myNumber
	fmt.Println("Value of ptr1 is", ptr1)

	*ptr1 = *ptr1 + 1
	fmt.Println("Updated myNumber is", myNumber)
}

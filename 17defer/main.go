// https://youtu.be/jiy584-vv38?si=ilRNG1Ib6zITPHuz
package main

import "fmt"

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v\n", i)
	}
}

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Welcome to defer")
	fmt.Println("Hello")
	myDefer()
}

// Defer works in LIFO manner.
// Meaning, the last declared 'Defer' will run first

package main

import "fmt"

func printCustom() {
	fmt.Println("PrintCustom")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	var total int = 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println("Functions in Golang")
	printCustom()
	fmt.Printf("Sum of %v and %v is %v\n", 2, 3, adder(2, 3))
	fmt.Printf("%v\n", proAdder(2, 4, 34, 123))
}

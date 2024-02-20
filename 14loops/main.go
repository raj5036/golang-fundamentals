package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	var days = []string{"Sunday", "Tuesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// For Loop
	for d := 0; d < len(days); d++ {
		fmt.Printf("Day %v: %v\n", d, days[d])
	}

	// Foreach Loop
	for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index, day)
	}

	// While Loop
	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 5 {
			break
		}
		fmt.Printf("rogueValue: %v\n", rogueValue)
		rogueValue++
	}
}

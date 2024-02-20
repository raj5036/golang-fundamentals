package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")

	var fruits = []string{"Apple", "Orange"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	fruits = append(fruits, "Guava")
	fruits = append(fruits, "Grapes", "Banana")
	fmt.Println("After appending", fruits)

	// Better way to use Slices
	scores := make([]int, 3)
	scores[0] = 1
	scores[1] = 2
	scores[2] = 3
	fmt.Println("scores", scores)

	scores = append(scores, 23, 12)
	fmt.Println("scores after append", scores)

	sort.Ints(scores)
	fmt.Println("scores after sort", scores)

	fmt.Println("scores => IsSorted", sort.IntsAreSorted(scores))
}

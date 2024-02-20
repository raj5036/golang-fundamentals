package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Walrus operator in golang: https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Pizza: ")

	// Comma OK || err OK syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating us ", input)
	fmt.Printf("Entered input is of type is %T", input)

}

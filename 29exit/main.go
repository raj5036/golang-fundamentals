// https://gobyexample.com/exit

package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Printing Statement")

	os.Exit(1)
}

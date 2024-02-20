package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	numRating, _ := reader.ReadString('\n')

	fmt.Printf("Your rating is %v", numRating)
	numRatingConv, err := strconv.ParseFloat(strings.TrimSpace(numRating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After adding 1 to your rating:", numRatingConv+1)
	}
}

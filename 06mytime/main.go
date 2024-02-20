package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of Golang")
	presentTime := time.Now()
	fmt.Println("presentTime", presentTime)
	fmt.Println("presentTime.format", presentTime.Format("01-02-2006"))
}

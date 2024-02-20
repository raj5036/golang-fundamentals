package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Handling Channels here")
	wg := &sync.WaitGroup{}
	// mut := &sync.RWMutex{}
	myCh := make(chan int, 1)

	wg.Add(2)
	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 3
		wg.Done()
	}(myCh, wg)
	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("Printing contents from myCh", <-myCh)
		fmt.Println("Printing contents from myCh", <-myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}

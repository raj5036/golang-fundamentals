package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Handling race condition with mutex")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	score := []int{}

	// Immediately Invoked Functions (as Goroutine)
	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait() // Asking main function to 'wait' until all goroutines have been executed
	fmt.Println(score)
}

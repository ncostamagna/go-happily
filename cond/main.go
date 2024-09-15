package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(3)

	go func() {
		fmt.Println("Goroutine 1 is started")
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 1 is waiting for condition")
		cond.Wait()
		fmt.Println("Goroutine 1 met the condition")

		fmt.Println("Goroutine 1 is done")
	}()

	go func() {
		fmt.Println("Goroutine 3 is started")
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 3 is waiting for condition")
		cond.Wait()
		fmt.Println("Goroutine 3 met the condition")

		fmt.Println("Goroutine 3 is done")
	}()

	go func() {
		fmt.Println("Goroutine 2 is started")
		defer wg.Done()

		time.Sleep(2 * time.Second) // Simulating some work

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 2 is signaling condition")
		//cond.Signal() // one goroutine
		cond.Broadcast() // all goroutines
		fmt.Println("Goroutine 2 completed signaling")

		fmt.Println("Goroutine 2 is done")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
}

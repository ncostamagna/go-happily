package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// atomic is faster than mutex but has more limitarion
	var counter int64 
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Atomic increment
		}()
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Atomic increment
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // Expected: 10





	counter = 0
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mu.Lock()   // Lock before modifying the shared resource
			counter++   // Critical section
			mu.Unlock() // Unlock after modifying
		}()

		go func() {
			defer wg.Done()
			mu.Lock()   // Lock before modifying the shared resource
			counter++   // Critical section
			mu.Unlock() // Unlock after modifying
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // Expected: 10
}

package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	fmt.Printf("start - count %d \n", c.v[key])

	c.mu.Lock()
	fmt.Printf("start lock - count %d \n", c.v[key])
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
	defer wg.Done()
	fmt.Printf("end - count %d \n", c.v[key])
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	var wg sync.WaitGroup
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("somekey", &wg)
	}

	wg.Wait()
	fmt.Println(c.Value("somekey"))
}

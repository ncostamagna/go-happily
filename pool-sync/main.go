package main

import (
	 "fmt"
    "math/rand"
	"sync"
	"time"
)

// another example
type resource struct {
	id int
	mutex sync.Mutex
}
func newResource() *resource {
	return &resource{
			id: rand.Intn(100), // Generate a random ID
	}
}
func process() {
	// Create a sync.Pool to manage resources
	var pool sync.Pool
	pool.New = func() interface{} {
		fmt.Println("Creating new resource with pool")
		return newResource()
	}

	nr := newResource()

	wait := sync.WaitGroup{}
	wait.Add(10)
	// Function to acquire a resource from the pool
	getResource := func() *resource {
		
		// when we get a resource from the pool, it's removed from the pool
		if v := pool.Get(); v != nil {
			fmt.Printf("Using resource from pool: %p\n", v)
			return v.(*resource)
		}

		fmt.Println("Using resource from new")
		// new value, new address
		return pool.New().(*resource)
	}

	// Simulate using resources
	for i := 0; i < 10; i++ {
		pool.Put(nr)
		go func() {
			defer wait.Done()
			r := getResource()
			fmt.Printf("Using resource with ID from pool: %d\n", r.id)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // Simulate work
		}()
	}
	wait.Wait()
}

func main() {
	process()
}

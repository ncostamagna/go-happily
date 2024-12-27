package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())
	bufPool.Put(b)
}

// another example
type resource struct {
	id int
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

	// Function to release a resource back to the pool
	releaseResource := func(r *resource) {
		pool.Put(r)
	}

	getWithoutPool := func() *resource {
		fmt.Println("Creating new resource without pool")
		return newResource()
	}
	// Simulate using resources
	for i := 0; i < 10; i++ {
		r := getResource()
		r2 := getWithoutPool()
		fmt.Printf("Using resource with ID from pool: %d\n", r.id)
		fmt.Printf("Using resource with ID from new: %d\n", r2.id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // Simulate work
		releaseResource(r)
	}
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")

	Log(os.Stdout, "path", "/search?q=flowers")

	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	process()
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=flowers")
}

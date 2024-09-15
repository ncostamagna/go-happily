package main

import (
    "sync"
	"fmt"
	"time"
	"math/rand"
)

func main() {
    loadMap()
	mapRange()
	mapConcurrency()
	cotidionalMapConcurrency()



	c := Cache{}

    // Launch multiple goroutines to access the cache
    for i := 0; i < 10; i++ {
        go accessCache(&c, i)
    }

    // Wait for all goroutines to finish
    time.Sleep(5 * time.Second)
}


func loadMap(){
	fmt.Println("###### Load Map ######")
	var m sync.Map

    // Adding key-value pairs concurrently
    m.Store("key1", "value1")
    m.Store("key2", "value2")

    // Retrieving value for a key
    if val, ok := m.Load("key1"); ok {
        println(val.(string)) // Output: value1
    }
	// func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
	// func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)

	if val, ok := m.Load("key2"); ok {
        println(val.(string)) // Output: value1
    }
    // Deleting a key
    m.Delete("key2")

	if val, ok := m.Load("key1"); ok {
        println(val.(string)) // Output: value1
    }

	if val, ok := m.Load("key2"); ok {
        println(val.(string)) // Output: value1
    }
	fmt.Println(m)
	fmt.Println("######################")
}

func mapRange() {
	fmt.Println("####### Range #######")
	var m sync.Map

    m.Store("key1", "value1")
    m.Store("key2", "value2")

    // Iterating over sync.Map
    m.Range(func(key, value interface{}) bool {
        println(key.(string), value.(string))
        return true
    })
	fmt.Println("#####################")
}

func mapConcurrency(){
	fmt.Println("###### Concurrency ######")
	var m sync.Map
    var wg sync.WaitGroup
	fmt.Println("first iteration")
    // Concurrent write operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
		time.Sleep(100 * time.Millisecond)
        go func(n int) {
            defer wg.Done()
            m.Store(n, n*10)
        }(i)
    }

    wg.Wait()
	fmt.Println("second iteration")
    // Concurrent read operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
		time.Sleep(100 * time.Millisecond)
        go func(n int) {
            defer wg.Done()
            if val, ok := m.Load(n); ok {
                println(val.(int))
            }
        }(i)
    }

    wg.Wait()
	fmt.Println("#########################")
}

func cotidionalMapConcurrency(){
	fmt.Println("###### Cotid Conc ######")
	m := make(map[int]int)
    var wg sync.WaitGroup
	fmt.Println("first iteration")
    // Concurrent write operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
		time.Sleep(100 * time.Millisecond)
        go func(n int) {
            defer wg.Done()
            m[n] = n*10
        }(i)
    }

    wg.Wait()
	fmt.Println("second iteration")
    // Concurrent read operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
		time.Sleep(100 * time.Millisecond)
        go func(n int) {
            defer wg.Done()
                println(m[n])
        }(i)
    }

    wg.Wait()
	fmt.Println("#########################")
}



// Cache struct using sync.Map
type Cache struct {
    store sync.Map
}

// Set a value in the cache
func (c *Cache) Set(key string, value interface{}) {
    c.store.Store(key, value)
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
    val, ok := c.store.Load(key)
    return val, ok
}

// Delete a value by key from the cache
func (c *Cache) Delete(key string) {
    c.store.Delete(key)
}

func accessCache(c *Cache, id int) {
    key := fmt.Sprintf("key%d", id)
    value := rand.Intn(100)

    // Set a value in the cache
    c.Set(key, value)
    fmt.Printf("Goroutine %d set %s to %d\n", id, key, value)

    // Get a value from the cache
    if val, ok := c.Get(key); ok {
        fmt.Printf("Goroutine %d got %s: %d\n", id, key, val)
    }

    // Sleep to simulate work
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

    // Delete the key
    c.Delete(key)
    fmt.Printf("Goroutine %d deleted %s\n", id, key)
}




/*
with sync.Maps is not necesarry to use Lock and Unlock
*/

func luckMap(){
	var mu sync.Mutex
    m := make(map[string]int)

    // Writing to the map
    mu.Lock()
    m["key1"] = 42
    mu.Unlock()

    // Reading from the map
    mu.Lock()
    value := m["key1"]
    mu.Unlock()

    fmt.Println("key1:", value)
}

func luckSyncMap(){
	var m sync.Map

    // Writing to the map
    m.Store("key1", 42)

    // Reading from the map
    if value, ok := m.Load("key1"); ok {
        fmt.Println("key1:", value)
    }

    // Deleting from the map
    m.Delete("key1")
}
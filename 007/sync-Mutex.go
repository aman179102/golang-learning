// sync.Mutex example – protecting shared data from concurrent access
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is a struct that stores a map and a mutex.
// The mutex is used to make sure that only one goroutine
// can access the map at a time.
//
// Without the mutex, simultaneous writes to the map
// would cause a race condition and crash the program.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for a given key.
// It locks the mutex before modifying the map,
// ensuring exclusive access.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()   // only ONE goroutine can pass this line at a time
	c.v[key]++    // safe update
	c.mu.Unlock() // release the lock so others can continue
}

// Value returns the value for a given key.
// We lock before reading, because reading and writing
// simultaneously is unsafe.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // ensure unlock even if function returns early
	return c.v[key]
}

func main() {
	// Initialize the SafeCounter with a proper map
	c := SafeCounter{v: make(map[string]int)}

	// Start 1000 goroutines that all increment the same key.
	// Without a mutex this would cause a data race.
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// Give goroutines time to finish (not ideal in production,
	// but fine for an example).
	time.Sleep(time.Second)

	// Print the final value – should be 1000.
	fmt.Println(c.Value("somekey"))
}

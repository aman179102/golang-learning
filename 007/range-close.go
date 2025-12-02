// range and close
package main

import (
	"fmt"
)

// fibonacci sends 'n' Fibonacci numbers into channel c.
// After sending all values, it closes the channel.
// Closing a channel signals that no more values will be sent.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Closing the channel tells the receiver that no more data is coming.
	close(c)
}

func main() {
	// Create a buffered channel with capacity 10.
	c := make(chan int, 10)

	// Start Fibonacci generator in another goroutine.
	go fibonacci(cap(c), c)

	// Using 'range' automatically receives values from the channel
	// until the channel is closed. No manual receive is needed.
	for i := range c {
		fmt.Println(i)
	}
}

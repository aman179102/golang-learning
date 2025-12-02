// buffered channels
package main

import "fmt"

func main() {
	// Create a channel with a buffer size of 2.
	// This means it can hold two values without waiting
	// for a receiver to read them.
	ch := make(chan int, 2)

	// These two sends happen immediately because the buffer
	// has enough space. No blocking occurs here.
	ch <- 1
	ch <- 2

	// Now we read the values back from the channel.
	// A buffered channel works in FIFO order.
	fmt.Println(<-ch) // prints 1
	fmt.Println(<-ch) // prints 2
}

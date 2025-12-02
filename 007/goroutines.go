package main

import (
	"fmt"
	"time"
)

// say prints a message 5 times with a small delay.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // just to slow it down
		fmt.Println(s)
	}
}

func main() {

	// Running say("world") in a goroutine.
	// This means it runs in the background, asynchronously.
	go say("world")

	// Running say("hello") normally.
	// This runs in the main goroutine.
	say("hello")
}

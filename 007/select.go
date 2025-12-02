// select with default case
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// tick sends a value every 100ms (like a ticker).
	tick := time.Tick(100 * time.Millisecond)

	// boom sends a value once after 500ms (like a timer).
	boom := time.After(500 * time.Millisecond)

	// Helper function to show time elapsed since start.
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		// This case runs every time 'tick' sends a value (every 100ms).
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())

		// This case triggers once when 500ms have passed.
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return

		// The default case runs when none of the channels are ready.
		// It prevents the select from blocking.
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond) // simulate doing other work
		}
	}
}

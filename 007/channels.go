package main

import "fmt"


// Channels are a typed conduit through which you can send and receive values with the channel operator, <-. 


// sum calculates the sum of a slice and sends the result into channel c
func sum(nums []int, c chan int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	c <- total // send result TO the channel
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}

	// create a channel that carries int values
	c := make(chan int)

	// run two goroutines: each calculates half of the list
	go sum(s[:len(s)/2], c) // first half
	go sum(s[len(s)/2:], c) // second half

	// receive two results from the channel
	x := <-c
	y := <-c

	fmt.Println("First Result:", x)
	fmt.Println("Second Result:", y)
	fmt.Println("Total:", x+y)
}

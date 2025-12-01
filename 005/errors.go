package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

////////////////////////////////////////////////////
// 1) Custom Error Example
////////////////////////////////////////////////////

// MyError stores time + message
type MyError struct {
	When time.Time
	What string
}

// Implement the Error() method to satisfy error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run() returns a custom error
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

////////////////////////////////////////////////////
// 2) Sqrt function with error handling
////////////////////////////////////////////////////

// Sqrt returns sqrt(x) or an error if x < 0
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot sqrt negative number")
	}
	return math.Sqrt(x), nil
}

////////////////////////////////////////////////////
// MAIN
////////////////////////////////////////////////////
func main() {

	// -------- Custom error example --------
	if err := run(); err != nil {
		fmt.Println("Custom Error:", err)
	}

	// -------- Sqrt function example --------
	result, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(2) =", result)
	}

	result, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

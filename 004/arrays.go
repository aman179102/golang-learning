package main

import (
	"fmt"
	"strings"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// array slicing
	//if you want array size actual to the size of the elements inside the array use [...]
	b := [...] int {2,3,4,5}
	fmt.Println(b[0:2])
	//slice literal
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	//nil slices
	var d []int
	fmt.Println(d, len(d), cap(d))
	if d == nil {
		fmt.Println("nil!")
	}
	//creating slice with make keyword
	x := make([]int,5)
	fmt.Println(x)
	//Slices of slices
	//Slice can contain any type, including other slices
		// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	//Appending to a slice
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	//range form of the for loop
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//You can skip the index or value by assigning to _
	
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


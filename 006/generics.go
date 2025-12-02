package main

import "fmt"

////////////////////////////////////////////////////
// 1) GENERIC INDEX FUNCTION
////////////////////////////////////////////////////

// Index returns the index of x in slice s, or -1 if not found.
// T must be a comparable type so we can use ==.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////
// 2) GENERIC LINKED LIST TYPE
////////////////////////////////////////////////////

// List represents a simple singly-linked list.
// T can be ANY type because we use `any` constraint.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new value at the *front* of the list.
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{next: l, val: v}
}

// Print prints all values in the list.
func (l *List[T]) Print() {
	for node := l; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}

////////////////////////////////////////////////////
// MAIN DEMO
////////////////////////////////////////////////////
func main() {

	// ---------- Example 1: Using Index[T] ----------
	ints := []int{10, 20, 15, -10}
	fmt.Println("Index of 15:", Index(ints, 15)) // 2

	strs := []string{"foo", "bar", "baz"}
	fmt.Println("Index of 'hello':", Index(strs, "hello")) // -1

	// ---------- Example 2: Using Generic Linked List ----------
	var list *List[int]
	list = list.Push(10)
	list = list.Push(20)
	list = list.Push(30)

	fmt.Println("\nPrinting int list:")
	list.Print()

	var slist *List[string]
	slist = slist.Push("apple")
	slist = slist.Push("banana")
	slist = slist.Push("cherry")

	fmt.Println("\nPrinting string list:")
	slist.Print()
}

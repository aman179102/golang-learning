package main

import "fmt"

////////////////////////////////////////////////////
// TYPE ASSERTION EXAMPLES
////////////////////////////////////////////////////
func typeAssertionDemo() {
	var i interface{} = "hello"

	// Direct assertion (must be correct type)
	s := i.(string)
	fmt.Println(s)

	// Safe assertion (returns value + ok)
	s, ok := i.(string)
	fmt.Println(s, ok)

	// Wrong type → ok = false
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Wrong type without ok → PANIC
	// Uncomment to see panic:
	// f = i.(float64)
	// fmt.Println(f)
}

////////////////////////////////////////////////////
// TYPE SWITCH EXAMPLES
////////////////////////////////////////////////////
func do(i interface{}) {
	switch v := i.(type) {

	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)

	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))

	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

////////////////////////////////////////////////////
// MAIN — runs both examples
////////////////////////////////////////////////////
func main() {

	fmt.Println("---- Type Assertion Demo ----")
	typeAssertionDemo()

	fmt.Println("---- Type Switch Demo ----")
	do(21)
	do("hello")
	do(true)
}

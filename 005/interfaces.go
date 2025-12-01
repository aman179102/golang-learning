package main

import (
	"fmt"
	"math"
)

////////////////////////////////////////////////////
// INTERFACE : "If you have M(), you belong here."
////////////////////////////////////////////////////
type I interface {
	M()
}

////////////////////////////////////////////////////
// TYPE 1 : T (has text inside)
////////////////////////////////////////////////////
type T struct {
	S string
}

// Value method (normal)
func (t T) M() {
	fmt.Println("Value T:", t.S)
}

// Pointer method (can handle nil)
func (t *T) M() {
	if t == nil {
		fmt.Println("Pointer T: <nil>")
		return
	}
	fmt.Println("Pointer T:", t.S)
}

////////////////////////////////////////////////////
// TYPE 2 : Float (numbers can also talk!)
////////////////////////////////////////////////////
type F float64

func (f F) M() {
	fmt.Println("Float:", f)
}

////////////////////////////////////////////////////
// EMPTY INTERFACE (can hold ANYTHING)
////////////////////////////////////////////////////
func show(x interface{}) {
	fmt.Printf("Value: %v   Type: %T\n", x, x)
}

////////////////////////////////////////////////////
// STRUCT WITH A METHOD (Abs example)
////////////////////////////////////////////////////
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

////////////////////////////////////////////////////
// MAIN PROGRAM (where everything runs)
////////////////////////////////////////////////////
func main() {

	// 1) Value receiver example
	var i I = T{"hello"}
	i.M()

	// 2) Pointer receiver example (nil safe)
	var p *T
	i = p // i stores nil pointer
	i.M()

	i = &T{"world"}
	i.M()

	// 3) Float implementing the interface
	i = F(3.14)
	i.M()

	// 4) Empty interface can hold anything
	show(nil)
	show(100)
	show("hi there")

	// 5) Struct + method example
	v := Vertex{3, 4}
	fmt.Println("Abs value:", v.Abs())
}

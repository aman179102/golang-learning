package main

import (
	"fmt"
	"math"
)

// A Point just holds two numbers: X and Y.
// You can think of it like a spot on a graph.
type Point struct {
	X, Y float64
}

// A small number type so we can add a helper method.
type Num float64

// Turns any number into its positive form.
func (n Num) Absolute() float64 {
	if n < 0 {
		return float64(-n)
	}
	return float64(n)
}

// Finds how far the point is from (0,0).
func (p Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Same distance calculation, just another name.
func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Makes the point larger or smaller.
// Both X and Y get multiplied.
func (p *Point) Multiply(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// A standalone version of Multiply.
func VectorScale(p *Point, f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// Same distance formula again.
func VectorLength(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Another way to scale the point.
func (p *Point) ScaleBy(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// The same scaling idea in a simple function.
func ScaleByFunc(p *Point, f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// Checks how big the point is from the origin.
func (p Point) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Function version of the same idea.
func MagnitudeFunc(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Stretches the point outward.
// Mostly same as scaling.
func (p *Point) Stretch(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// Another distance method from a pointer.
func (p *Point) Size() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	// Basic length example
	a := Point{3, 4}
	fmt.Println(a.Length())

	// Positive number example
	n := Num(-math.Sqrt2)
	fmt.Println(n.Absolute())

	// Multiply demo
	b := Point{3, 4}
	b.Multiply(10)
	fmt.Println(b.Distance())

	// Using VectorScale
	c := Point{3, 4}
	VectorScale(&c, 10)
	fmt.Println(VectorLength(c))

	// Using ScaleBy and then the function version
	d := Point{3, 4}
	d.ScaleBy(2)
	ScaleByFunc(&d, 10)

	// Pointer example
	e := &Point{4, 3}
	e.ScaleBy(3)
	ScaleByFunc(e, 8)
	fmt.Println(d, e)

	// Magnitude tests
	f := Point{3, 4}
	fmt.Println(f.Magnitude())
	fmt.Println(MagnitudeFunc(f))

	// Same thing with a pointer
	g := &Point{4, 3}
	fmt.Println(g.Magnitude())
	fmt.Println(MagnitudeFunc(*g))

	// The latest part you added
	h := &Point{3, 4}
	fmt.Printf("Before scaling: %+v, Size: %v\n", h, h.Size())
	h.Stretch(5)
	fmt.Printf("After scaling: %+v, Size: %v\n", h, h.Size())
}

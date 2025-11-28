package main

// -------------------------------
// IMPORTS (bringing tools)
// -------------------------------
import (
	"fmt"      // for printing
	"math"     // for math functions
	"math/rand"// for random numbers
	"time"     // for current time
)

// -------------------------------
// CONSTANTS & VARIABLES
// -------------------------------

// A constant never changes
const Pi = 3.14

// Variables declared outside the function
var rahul bool
var k, l = true, false

func main() {

	// -------------------------------
	// BASIC PRINTING
	// -------------------------------

	fmt.Println("The current time is:", time.Now())
	fmt.Println("My favourite number is:", rand.Intn(10))
	fmt.Println("The square root of 25 is:", math.Sqrt(25))
	fmt.Println("The value of built-in Pi is:", math.Pi)

	// -------------------------------
	// FUNCTIONS
	// -------------------------------

	fmt.Println("5 + 5 =", add(5, 5))
	fmt.Println("5 * 5 =", mul(5, 5))

	// Returning multiple values
	a, b := swap("Aman", "Kumar")
	fmt.Println("Swapped:", a, b)

	// Automatic splitting example
	fmt.Println("Split 17:", split(17))

	// -------------------------------
	// VARIABLES
	// -------------------------------

	var kirti int // default value (0)
	fmt.Println("Marks → Kirti:", kirti, "Rahul:", rahul)

	// Variables with initial values
	var i, j = 1, 2
	fmt.Println(i, j, k, l)

	// Short variable declaration (only inside functions)
	ak := true
	fmt.Println("Short variable ak:", ak)

	// -------------------------------
	// ZERO VALUES
	// -------------------------------

	var e int
	var f float64
	var w bool
	var s string
	fmt.Printf("%v %v %v %q\n", e, f, w, s)

	// -------------------------------
	// TYPE CONVERSION
	// -------------------------------

	var x, y int = 3, 4
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(g)
	fmt.Println("Triangle values:", x, y, z)

	// -------------------------------
	// TYPE INFERENCE
	// -------------------------------

	v := 42 // Go guesses this is an int
	fmt.Printf("v is of type %T\n", v)

	// -------------------------------
	// CONSTANTS
	// -------------------------------

	fmt.Println("Pi constant:", Pi)

	// -------------------------------
	// NUMERIC CONSTANTS
	// (Big scary numbers made easy)
	// -------------------------------

	// Go can handle very big numbers safely
	fmt.Println("needInt(Small):", needInt(Small))
	fmt.Println("needFloat(Small):", needFloat(Small))
	fmt.Println("needFloat(Big):", needFloat(Big))
}

// -------------------------------
// FUNCTIONS
// -------------------------------

// Simple addition
func add(x int, y int) int {
	return x + y
}

// Simple multiplication
func mul(x, y int) int {
	return x * y
}

// Return two values (swapping)
func swap(x, y string) (string, string) {
	return y, x
}

// Returns two numbers from one input
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// -------------------------------
// BIG NUMBERS (Numeric Constants)
// -------------------------------

const (
	// 1 shifted left 100 times → a huge number
	Big = 1 << 100

	// Shift right 99 times → becomes 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

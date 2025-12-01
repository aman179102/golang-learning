package main

import (
	"fmt"
	"strconv"
)

////////////////////////////////////////////////////
// 1) Stringer Example using Person struct
////////////////////////////////////////////////////

// Person will print in a nice human-readable format
type Person struct {
	Name string
	Age  int
}

// String() method → custom printing
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

////////////////////////////////////////////////////
// 2) IPAddr Stringer Exercise
////////////////////////////////////////////////////

// IPAddr is just 4 bytes, like 127.0.0.1
type IPAddr [4]byte

// String() method → convert [4]byte to "x.x.x.x"
func (ip IPAddr) String() string {
	// convert each byte to string, join with dots
	return strconv.Itoa(int(ip[0])) + "." +
		strconv.Itoa(int(ip[1])) + "." +
		strconv.Itoa(int(ip[2])) + "." +
		strconv.Itoa(int(ip[3]))
}

////////////////////////////////////////////////////
// MAIN — runs both demos
////////////////////////////////////////////////////
func main() {

	// -------- Person Stringer Example --------
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println("People:")
	fmt.Println(a)
	fmt.Println(z)

	// -------- IP Address Stringer Example --------
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	fmt.Println("\nIP Addresses:")
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

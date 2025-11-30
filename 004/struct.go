package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	//struct values are accessed using a dot 
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	//struct values can be accessed through a struct pointer
	p := &v
	p.X = 1e9  // this 1e9 is also called 1*10^9
	fmt.Println(v)
	//struct literals 
	fmt.Println(v1,p,v2,v3)
}

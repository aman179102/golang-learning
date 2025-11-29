package main

import (
	"fmt"
	
)


func main(){
	
	//defer is executed at the end of the function no matter which statement is in the defer
	defer fmt.Println("Golang")
	fmt.Print("Hello from ")
	//we can stack defer statements
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	
}
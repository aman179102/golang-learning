package main

import (
	"fmt"
	"runtime"
	"time"
)


func main(){
	
	//**************
	//if************
	//**************
	if 27<50 {
		fmt.Println("27 is smaller than 50")
	}
	//You can run short statement in if after the condition checking 
	if fmt.Println("Hello from if statement"); 1<2{
		fmt.Println("2 is greater than 1")
	} else {
		fmt.Println("This is not a world ")
	}
	//variables declared in short if statement can also be accessible in the else statement
	//**************
	//Switch********
	//**************
		fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	//break keyword is not necessary in golang 
	//switch with no condition 
	// switch with no condition is set to true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
}
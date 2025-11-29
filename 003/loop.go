package main

import "fmt"


//golang language has only one loop called for loop


func main(){
	sum := 0
	for i := 0; i<10; i++ {
		sum+=i
	} 
	fmt.Println("The sum from 0 to 9 is",sum)
	//**********************
	//init and post statements are optional
	//**********************
	summation := 0
	for ; summation<1000; {
		summation+=1
	}
	fmt.Println("Final result",summation)
	//**********************************
	//We can use for as while loop *****
	//**********************************
	icon := 0
	for icon<10 {
		fmt.Print(icon," ")
		icon++
	}
	fmt.Print("\n")
	//********************
	//Infinte Loop *******
	//********************
	
	//for {
	//
	//}

}
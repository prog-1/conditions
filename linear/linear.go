package main

import (
	"fmt"
)

func main(){
	var a, b, r float64
	fmt.Println("The program solves a linear inequality Ax + B > 0.Enter A and B: ")
	fmt.Scan(&a, &b)
	if a < 0 {
		x := b / a
		fmt.Println("x <", x)
	}
	
	else if a > 0{
		x := -b / a
		fmt.Prinln("x >", x)
	}
	
	else {
		fmt.Println("result can't be found")
	} 

}
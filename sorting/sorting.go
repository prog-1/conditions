package main

import (
	"fmt"
)

func main(){
	var a, b, c float64
	fmt.Println("The program sorts three numbers.Enter the numbers: ")
	fmt.Scan(&a,&b,&c)
    if a > c, a > b ,c > b {
		a, b, c = a, c, b
	}
	else if c > b, c > a, b > a {
		a, b, c = c, b, a
	}
	else if b > a, b > c, a > c {
		a, b, c = b, a, c
	}
	else if a > b, b > c, a > c {
		a, b, c = a, b, c
	}
	else if b > c, b > a, c > a {
		a, b, c = b, c, a
	}
	else if c > a, c > b, a > b {
		a, b, c = c, a, b
	}
	fmt.Println("Numbers in a sorted order: ", a, b, c)

}
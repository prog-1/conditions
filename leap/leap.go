package main

import (
	"fmt"
)

func main(){
	var g float64
	fmt.Println("The program determines if a year is a leap year.Enter the year: ")
	fmt.Scan(&g)
	if g%4 == 0 {
		fmt.Println(g," is a leap year")
	} else if g%100 == 0 {
		if g%400 == 0{
			fmt.Println(g," is a leap year")
		} else {
			fmt.Println(g," isn't a leap year")
		}
	} else {
		fmt.Println(g," isn't a leap year")
	}
}
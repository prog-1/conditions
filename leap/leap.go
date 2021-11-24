package main

import (
	"fmt"
)

func main(){
	var g int
	fmt.Println("The program determines if a year is a leap year.Enter the year: ")
	fmt.Scan(&g)
	if g%4 == 0 {
		fmt.Println(g," is a leap year")
	} else if g%100 == 0 {
		if g%400 == 0{ //1900 не делится на 400,но почему показывает что да?
			fmt.Println(g," is a leap year")
		} else {
			fmt.Println(g," isn't a leap year")
		}
	} else {
		fmt.Println(g," isn't a leap year")
	}
}
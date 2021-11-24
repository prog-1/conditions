package main

import (
	"fmt"
	"math"
)

func main(){
	var a, b, c float64
	fmt.Println("The program finds solutions of quadratic equation.Enter the coefficients A, B and C: ")
	fmt.Scan(&a,&b,&c)
	d := b * b - 4 * a * c
	if d > 0 {
		x1 := (-b - math.Sqrt(d)) / 2 * a 
		x2 := (-b + math.Sqrt(d)) / 2 * a 
		fmt.Println("x1= ",x1,"x2= ",x2)
	} else if d == 0 {
		x := -d / 2 * a
		fmt.Println("x= ",x)
	} else if d < 0 {
		fmt.Println("no roots")
	}
}
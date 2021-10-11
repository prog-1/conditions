package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(`The program sorts three numbers.
	Enter the numbers:`)
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	max2 := (a + b + math.Abs(a-b)) / 2
	max := (c + max2 + math.Abs(c-max2)) / 2
	min2 := (a + b - math.Abs(a-b)) / 2
	min := (c + min2 - math.Abs(c-min2)) / 2
	if a != min && a != max {
		fmt.Print(min, a, max)
	} else {
		if b != min && b != max {
			fmt.Print(min, b, max)
		} else {
			fmt.Print(min, c, max)
		}
	}

}

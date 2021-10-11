package main

import "fmt"

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
    (1) if you want to convert km/h to m/s;
    (2) if you want to convert m/s to km/h.`)
	var a float64
	fmt.Scanln(&a)
	if a == 1 {
		fmt.Println("Enter speed km/h")
		var v float64
		fmt.Scanln(&v)
		ms := v / 3.6
		fmt.Println(v, "km/h =", ms, "m/s")
	}
	if a == 2 {
		fmt.Println("Enter speed m/s")
		var b float64
		fmt.Scanln(&b)
		km := b * 3.6
		fmt.Println(b, "m/s =", km, "km/h")
	} else {
		fmt.Println("Wrong number")
	}
}

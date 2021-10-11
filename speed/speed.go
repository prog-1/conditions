package main

import "fmt"

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
(1) if you want to convert km/h to m/s;
(2) if you want to convert m/s to km/h.`)
	var num float64
	fmt.Scan(&num)
	if num == 1 {
		fmt.Println("Enter the speed in km/h")
		var k float64
		fmt.Scan(&k)
		fmt.Println("The speed in m/s is", k/3.6)
	} else if num == 2 {
		fmt.Println("Enter the speed in m/s")
		var m float64
		fmt.Scan(&m)
		fmt.Println("The spped in km/h is", m*3.6)
	}
}

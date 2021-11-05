package main

import "fmt"

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h.`)
	var s float64
	fmt.Scanln(&s)
	if s == 1 {
		fmt.Println("Speed in km/h:")
		var km float64
		fmt.Scanln(&km)
		m := km / 3.6
		fmt.Println("Speed in ms:", m)
		if s == 2 {
			fmt.Println("Speed in m/s:")
			var m float64
			fmt.Scanln(&m)
			k := m * 3.6
			fmt.Println("Speed in kmh:", k)
		}
	}
}

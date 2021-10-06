package main

import "fmt"

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
		(1) if you want to convert km/h to m/s;
		(2) if you want to convert m/s to km/h`)
	var choose float64
	fmt.Scan(&choose)
	if choose == 1 {
		fmt.Println("Enter the speed in km/h")
		var kmh float64
		fmt.Scan(&kmh)
		fmt.Println("The speed in m/s is", kmh/3.6)
	} else if choose == 2 {
		fmt.Println("Enter the speed in m/s")
		var ms float64
		fmt.Scan(&ms)
		fmt.Println("The speed in km/h is", ms*3.6)
	}

}

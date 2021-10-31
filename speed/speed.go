package main

import (
	"fmt"
)

func main(){
	var n,km,m float64
	fmt.Println("This program performs km/h <-> m/s speed conversion.
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h. 
	Please enter: ")
	fmt.Scan(&n)
	if n == 1{
		//km to ms
		fmt.Println("Enter the speed in km/h: ")
		fmt.Scan(&km)
		s1 := 360 / (km * 1000)	
		fmt.Println(s1)
	}

	if n == 2{
		//ms to km
		fmt.Println("Enter the speed in m/s: ")
		fmt.Scan(&m)
		s2 := (1000 * m) / 360	
		fmt.Println(s2)
	}
}
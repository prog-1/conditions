package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter three triangle sides:")
	var x, y, z float64
	fmt.Scan(&x, &y, &z)
	if x+y > z && y+z > x && z+x > y {
		s := (x + y + z) / 2
		area := math.Sqrt(s * (s - x) * (s - y) * (s - z))
		fmt.Println("Triangle area:", area)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}

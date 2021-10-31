package main

import (
	"fmt"
)

func main(){
	var a,b,p,q,r float64
	fmt.Println("Enter the hole dimensions (2 numbers): ")
	fmt.Scan(&a,&b)
	fmt.Println("Enter the brick dimensions (3 numbers): ")
	fmt.Scan(&p,&q,&r)
	if a >= p || a >= q || a >= r || b >= p || b >= q || b >= r {
		fmt.Println("The brick can be inserted in the hole!")
	}  else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
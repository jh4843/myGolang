package main

import (
	"fmt"
)


func main() {
	var t [5]float64 = [5]float64{2.3, 51.2, 23.1, 22.3, 1.2}

	for i:=0; i< 5; i++ {
		fmt.Println(t[i]);	
	}

	var s = [5]int{1:10, 3:30}	// index:value
	for i:=0; i< 5; i++ {
		fmt.Println(s[i]);	
	}

	x := [...]int{10, 20, 30} // fixed size array = [3]int
	y := []int{10, 20, 30} // slice (not fixed array)
	for i:=0; i< len(x); i++ {
		fmt.Println(x[i]);	
	}

	for i:=0; i< len(y); i++ {
		fmt.Println(y[i]);	
	}

}
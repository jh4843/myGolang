package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {

	fmt.Printf("%0.18f", math.Nextafter(a, b))

	if(a == b) {
		fmt.Println("equal !!")
	} else {
		fmt.Println("diff !!")
	}

	return math.Nextafter(a, b) == b
}

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2)

	fmt.Printf("x:%08b x>>2:%08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2:%08b y>>2: %d\n", y, y>>2, y>>2)

	var a = 1.0
	var b = 2.0
	var c = 3.0


	fmt.Println(equal(a+b, c))

}
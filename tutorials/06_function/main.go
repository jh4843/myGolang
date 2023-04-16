package main

import (
	"fmt"
)

func Add(a int, b int) int {
	return a+b
}

// multi return
func Divide(a, b int) (int, bool) {
	if (b==0) {
		return 0, false
	}

	return a / b, true
}

// multi return with argument name
func DivideEx(a, b int) (res int, isSuccess bool) {
	if (b==0) {
		res = 0
		isSuccess = false
	}

	res = a/b
	isSuccess = true

	return 
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)

	res, isSuccess := Divide(9, 3)

	fmt.Printf("%d %t \n", res, isSuccess)

	res, isSuccess = DivideEx(9, 3)

	fmt.Printf("%d %t \n", res, isSuccess)
}
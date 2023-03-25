package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// out
	var a int = 1234
	var b int = 2567
	var c int = 256789123
	var f float64 = 322799.9787
	
	fmt.Print("a:", a, "b:", b, "f:", f);	// there is no space between ,
	fmt.Println("a:", a, "b:", b, "f:", f); // there is space between , & add \n at last.
	fmt.Printf("a: %d b: %d c: %f\n\n", a, b, f); //

	fmt.Printf("%5d, %5d, %5d\n", a, b, c)
	fmt.Printf("%05d, %05d, %05d\n", a, b, c)
	fmt.Printf("%-5d, %-5d, %-5d\n", a, b, c)

	fmt.Println();

	// Input
	var str string = ""
	var strLine string = ""

	fmt.Scan(&str, &strLine)
	fmt.Printf("scan: %s\n", str);
	fmt.Printf("scan: %s\n", strLine);


	// File handling
	stdin := bufio.NewReader(os.Stdin);

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println("err: ", err);
		strTemp, errTemp := stdin.ReadString('\n'); // \n 만날때까지 읽어라.

		fmt.Println(strTemp, errTemp)
	} else {
		fmt.Println(a, b, "count: ", n);
	}
}


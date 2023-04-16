package main

import (
	"fmt"
	"tutorials/09_module/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"

	"tutorials/09_module/exinit"
)

func main() {
	custompkg.PrintCustom()
	//custompkg.printCustom1()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 9, 10, 11, 13}
	graph := asciigraph.Plot(data)
	
	fmt.Println(graph)	

	exinit.PrintD()
}
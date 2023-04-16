package main

import (
	"fmt"
)

type Data struct {
	value int
	data [200]int
}

func changeData(arg* Data) {
	arg.value = 999
	arg.data[100] = 999
}


func main() {
	var data Data

	changeData(&data)
	fmt.Printf("value = %d\n", data.value);
	fmt.Printf("data[100] = %d\n", data.data[100]);

	var p *Data = &Data{}
	var t *Data = &data
	k := &Data{}
	var p2 = new(Data)

	fmt.Printf("data = %d %d\n", &data, data.value);
	fmt.Printf("p = %d %d\n", &p, p.value);
	fmt.Printf("t = %d %d\n", &t, t.value);
	fmt.Printf("k = %d %d\n", &k, k.value);
	fmt.Printf("p2= %d %d\n", &p2, p2.value);
}
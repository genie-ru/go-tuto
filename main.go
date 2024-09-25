package main

import "fmt"

func main() {

	
	var (
		i int = 1
		s string = "hello"
		f64 float64 = 1.0
		t, f bool = true, false
	)
	fmt.Println(i, s, f64, t, f)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

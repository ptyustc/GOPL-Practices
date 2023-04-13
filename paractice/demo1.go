package main

import "fmt"

type FU float64

func main() {
	var f float64
	var fu FU
	// fu = f // cannot use f (variable of type float64) as FU value in assignment
	fu = 0
	// fmt.Println(f == fu) // cannot compare f == fu (mismatched types float64 and FU)
	fmt.Println(fu == 0) // true
	fmt.Println(f == float64(fu))
}

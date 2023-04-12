package main

import (
	"fmt"
	"godemo/tempconv"
)

func main() {
	fmt.Println(tempconv.Celsius.CToF(100))
	fmt.Println(tempconv.Kelvin.KToC(100))
}

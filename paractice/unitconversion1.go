package main

import (
	"fmt"
	"godemo/unitconversion"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		para, _ := strconv.ParseFloat(arg, 64)
		fmt.Println(unitconversion.Inch.IToM(unitconversion.Inch(para)))
		fmt.Println(unitconversion.Pound.PToKG(unitconversion.Pound(para)))
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari float64
	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari)
	fmt.Printf("\nBola dengan jejari %v memiliki volume %v dan luas kulit %v", jejari, (4*math.Pi*math.Pow(jejari, 3))/3, 4*math.Pi*math.Pow(jejari, 2))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var k, hasil float64 = 0, 1
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	for i := 0; i < int(k); i++ {
		hasil *= math.Pow(4*float64(i)+2, 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}
	fmt.Printf("Nilai akar 2 = %.10f", hasil)
}

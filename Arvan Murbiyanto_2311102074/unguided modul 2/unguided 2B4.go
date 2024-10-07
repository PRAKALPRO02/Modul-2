package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var hasil float64 = 1.0

	for {
		fmt.Print("Nilai K = ")
		fmt.Scan(&k)
		if k < 0 {
			break
		}
		hasil *= math.Pow(4*float64(k)+2, 2) /
			((4*float64(k) + 1) * (4*float64(k) + 3))
		fmt.Printf("Nilai Akar 2 = %.10f \n", hasil)
	}
}
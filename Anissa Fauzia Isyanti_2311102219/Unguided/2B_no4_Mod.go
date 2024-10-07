package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	hasil := 1.0

	for i := 0; i <= k; i++ {
		hasil *= math.Pow(4*float64(i)+2, 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}

	fmt.Printf("\nNilai Akar 2 = %.10f \n", hasil)
}

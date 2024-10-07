package main

import (
	"fmt"
	"math"
)

func calculateResult(k int) float64 {
	result := 1.0
	for i := 0; i <= k; i++ {
		numerator := math.Pow(4*float64(i)+2, 2)
		denominator := (4*float64(i) + 1) * (4*float64(i) + 3)
		result *= numerator / denominator
	}
	return result
}

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	hasil := calculateResult(k)
	fmt.Printf("Nilai Akar 2 = %.10f \n", hasil)
}

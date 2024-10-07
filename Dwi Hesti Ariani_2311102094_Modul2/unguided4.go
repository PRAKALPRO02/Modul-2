package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Print("Masukkan nilai k: ")
	fmt.Scanln(&k)

	result := calculateSqrt2(k)
	fmt.Printf("Akar 2 untuk k = %d adalah %.10f\n", k, result)
}

func calculateSqrt2(k int) float64 {
	product := 1.0
	for i := 0; i <= k; i++ {
		numerator := math.Pow(float64(4*i+2), 2)
		denominator := float64((4*i + 1) * (4*i + 3))
		product *= numerator / denominator
	}
	return product
}

// Dwi Hesti Ariani_2311102094

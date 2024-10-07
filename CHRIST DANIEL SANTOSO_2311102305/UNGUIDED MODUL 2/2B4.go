package main

import (
	"fmt"
	"math"
)

func main() {
	var nilaiK int

	fmt.Print("Masukkan nilai k: ")
	fmt.Scan(&nilaiK)

	numerator := calculateNumerator(nilaiK)
	denominator := calculateDenominator(nilaiK)

	fk := float64(numerator) / float64(denominator)
	fmt.Printf("f(%d) = %.10f\n", nilaiK, fk)

	akarDua := math.Sqrt(2)
	fmt.Printf("Akar 2 = %.10f\n", akarDua)
}

func calculateNumerator(k int) int {
	return (4*k + 2) * (4*k + 2)
}
func calculateDenominator(k int) int {
	return (4*k + 1) * (4*k + 3)
}

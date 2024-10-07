package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	fmt.Print("Masukkan nilai k: ")
	fmt.Scan(&k)

	numerator := (4*k + 2) * (4*k + 2)
	denominator := (4*k + 1) * (4*k + 3)

	fk := float64(numerator) / float64(denominator)
	fmt.Printf("f(%d) = %.10f\n", k, fk)

	akar2 := math.Sqrt(2)

	fmt.Printf("Akar 2 = %.10f\n", akar2)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	result := 1.0
	for k := 0; k <= K; k++ {
		numerator := math.Pow((4*float64(k) + 2), 2)
		denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
		result *= numerator / denominator
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}

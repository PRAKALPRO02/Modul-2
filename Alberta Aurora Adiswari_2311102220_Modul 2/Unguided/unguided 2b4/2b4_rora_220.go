package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	numerator_220 := math.Pow((4*k + 2), 2)
	denominator_220 := (4*k + 1) * (4*k + 3)
	fk := numerator_220 / denominator_220

	fmt.Printf("Nilai f(K) = %.10f\n", fk)
}

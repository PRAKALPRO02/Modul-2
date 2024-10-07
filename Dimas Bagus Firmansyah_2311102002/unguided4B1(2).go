package main

import (
	"fmt"
	"math/big"
)
func calculateSqrt2(K int) *big.Float {
	
	result := big.NewFloat(1.0)

	for k := 0; k < K; k++ {
		
		numerator := float64((4*k + 2) * (4*k + 2))
		denominator := float64((4*k + 1) * (4*k + 3))

		term := big.NewFloat(numerator / denominator)
		result.Mul(result, term)
	}

	return result
}

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)
	result := calculateSqrt2(K)
	resultStr := fmt.Sprintf("%.10f", result)
	fmt.Println("Nilai akar 2 =", resultStr)
}

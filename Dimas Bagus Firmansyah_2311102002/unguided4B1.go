package main

import (
	"fmt"
	"math"
)

func main() {
	var K float64

	// Meminta input dari pengguna untuk nilai K
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Menghitung f(K) berdasarkan persamaan
	numerator := math.Pow(4*K+2, 2)
	denominator := (4*K + 1) * (4*K + 3)
	fK := numerator / denominator

	// Menampilkan hasil perhitungan
	fmt.Printf("Nilai f(K) = %.10f\n", fK)
}

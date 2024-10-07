package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(K)
func f(K int) float64 {
	numerator := math.Pow(float64(4*K+2), 2)
	denominator := float64((4*K + 1) * (4*K + 3))
	return numerator / denominator
}

// Fungsi untuk menghampiri sqrt(2) menggunakan rumus yang diberikan
func approximateSqrt2(K int) float64 {
	product := 1.0
	for k := 0; k <= K; k++ {
		product *= f(k)
	}
	return product
}

func main() {
	var K int

	// Bagian pertama: Menghitung f(K)
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)
	fmt.Printf("Nilai f(K) = %.10f\n", f(K))

	// Bagian kedua: Menghitung hampiran sqrt(2)
	fmt.Print("Masukkan nilai K untuk menghitung akar 2: ")
	fmt.Scan(&K)
	fmt.Printf("Nilai akar 2 = %.10f\n", approximateSqrt2(K))
}

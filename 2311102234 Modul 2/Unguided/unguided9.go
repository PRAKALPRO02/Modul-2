package main

import (
	"fmt"
	"math"
)

func main() {
	// Meminta input untuk nilai K dan menghitung f(K)
	var k float64
	fmt.Print("Masukkan nilai K untuk f(K): ")
	fmt.Scanln(&k)
	fK := calculateFK(k)
	fmt.Printf("Nilai f(K) = %.10f\n", fK)

	// Meminta input untuk nilai K dan menghitung akar2
	var K int
	fmt.Print("Masukkan nilai K untuk akar2: ")
	fmt.Scanln(&K)
	akar2 := calculateAkar2(K)
	fmt.Printf("Nilai akar2 = %.10f\n", akar2)
	fmt.Printf("Nilai akar 2 = %.10f\n", math.Sqrt(2))
}

// Fungsi untuk menghitung f(K)
func calculateFK(k float64) float64 {
	return (4 * k) / (k + k + 3)
}

// Fungsi untuk menghitung akar2
func calculateAkar2(K int) float64 {
	return (1 + float64(K) + 3) / (1 + 2)
}

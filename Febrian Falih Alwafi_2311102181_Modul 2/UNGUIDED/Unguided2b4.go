package main

import (
	"fmt"
	"math"
)

func main() {
	var K int

	// Minta pengguna untuk memasukkan nilai K
	fmt.Print("Nilai K= ")
	fmt.Scanln(&K)

	// Hitung akar2 dengan formula yang diberikan
	akar2 := calculateAkar2(K)

	// Tampilkan hasil
	fmt.Printf("Nilai akar2 = %.10f\n", akar2)
	fmt.Printf("Nilai akar 2 = %.10f\n", math.Sqrt(2))
}

// Fungsi untuk menghitung nilai akar2
func calculateAkar2(K int) float64 {
	return (1 + float64(K) + 3) / (1 + 2)
}

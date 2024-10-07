package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai f(k)
func hitungF(k int) float64 {
	pembilang := math.Pow(float64(4*k+2), 2)
	penyebut := float64((4*k + 1) * (4*k + 3))
	return pembilang / penyebut
}

func main() {
	var k int

	// Input nilai k
	fmt.Print("Nilai k = ")
	fmt.Scan(&k)

	// Hitung dan tampilkan nilai f(k)
	hasil := hitungF(k)
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)

	// Hitung hampiran akar 2
	akarDua := 1.0
	for i := 0; i <= k; i++ {
		akarDua *= hitungF(i)
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", akarDua)
}

//Zahrina Antika Malahati_2311102109

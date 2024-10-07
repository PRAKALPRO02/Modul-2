package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai f(k)
func hitungF(k int) float64 {
	// Menghitung pembilang dan penyebut dari rumus f(k)
	pembilang := math.Pow(float64(4*k+2), 2)
	penyebut := float64((4*k+1) * (4*k+3))
	return pembilang / penyebut
}

// Fungsi untuk menghitung perkalian f(i) hingga k untuk mendekati akar 2
func hitungAkarDua(k int) float64 {
	hasil := 1.0
	for i := 0; i <= k; i++ {
		hasil *= hitungF(i)
	}
	return hasil
}

func main() {
	// Input nilai k
	var k int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Menghitung nilai f(k)
	nilaiF := hitungF(k)
	fmt.Printf("Nilai f(%d) = %.10f\n", k, nilaiF)

	// Menghitung nilai aproksimasi akar 2 untuk k
	akarDua := hitungAkarDua(k)
	fmt.Printf("Nilai akar 2 untuk K = %d: %.10f\n", k, akarDua)
}

package main

import (
	"fmt"
	"math"
)

func hitungAkar2(K int) float64 {
	hasil := 1.0
	for k := 0; k < K; k++ {
		pembilang := math.Pow(4*float64(k)+2, 2)
		penyebut := (4*float64(k) + 1) * (4*float64(k) + 3)
		hasil *= pembilang / penyebut
	}
	return hasil
}

func main() {
	var K int
	fmt.Print("Masukkan nilai K = ")
	fmt.Scan(&K)

	akar2Hampiran := hitungAkar2(K)
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", akar2Hampiran)
}

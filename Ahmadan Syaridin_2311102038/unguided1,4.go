package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)
	hampiranAkar2 := 1.0
	for k := 0; k <= K; k++ {
		pembilang := (4*float64(k) + 2) * (4*float64(k) + 2)
		penyebut := (4*float64(k) + 1) * (4*float64(k) + 3)
		hampiranAkar2 *= pembilang / penyebut
	}
	fmt.Printf("Nilai hampiran √2 untuk K = %d adalah: %.10f\n", K, hampiranAkar2)
}

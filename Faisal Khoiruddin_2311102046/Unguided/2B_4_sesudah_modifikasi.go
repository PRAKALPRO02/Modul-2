package main

import (
	"fmt"
	"math"
)

func main() {
	var K float64
	fmt.Print("Nilai K: ")
	fmt.Scan(&K)

	akar2 := 1.0
	for k := 0; k <= int(K); k++ {
		pembilang := math.Pow(4*float64(k)+2, 2)
		penyebut := (4*float64(k) + 1) * (4*float64(k) + 3)
		akar2 *= pembilang / penyebut
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}

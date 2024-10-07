package main

import (
	"fmt"
	"math"
)

func main() {
	var K float64
	fmt.Printf("Nilai K: ")
	fmt.Scan(&K)

	pembilang := math.Pow(4*K+2, 2)
	penyebut := (4*K + 1) * (4*K + 3)
	fK := pembilang / penyebut

	fmt.Printf("Nilai f(K): %.10f\n", fK)
}

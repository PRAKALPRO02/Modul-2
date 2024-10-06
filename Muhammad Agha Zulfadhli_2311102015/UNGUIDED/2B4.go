package main

import (
	"fmt"
	"math"
)

func main() {
	var k, hasil float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	hasil = math.Pow(4*k+2, 2) / ((4*k + 1) * (4*k + 3))
	fmt.Printf("Nilai f(K) = %.10f", hasil)
}

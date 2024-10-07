package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var fungsi float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	pembilang := math.Pow(4*float64(k)+2, 2)
	penyebut := ((4.0*float64(k) + 1) * (4.0*float64(k) + 3))
	fungsi = pembilang / penyebut
	fmt.Printf("nilai f(k) = %.10f", fungsi)
}

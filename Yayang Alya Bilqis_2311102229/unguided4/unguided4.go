package main

import (
	"fmt"
)

func f(k int) float64 {
	return float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
}

func akar2(k int) float64 {
	var hasil float64 = 1
	for i := 0; i <= k; i++ {
		hasil *= f(i)
	}
	return hasil
}

func main() {
	var k int
	fmt.Print("Masukkan nilai k: ")
	fmt.Scan(&k)

	fmt.Printf("Nilai f(%d) = %.10f\n", k, f(k))
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2(k))
}

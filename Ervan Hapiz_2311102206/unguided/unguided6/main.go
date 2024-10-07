package main

import (
	"fmt"
)

func main() {
	var k int
	var F float64

	fmt.Print("Nilai k = ")
	fmt.Scanln(&k)

	F = float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	fmt.Printf("Nilai f(k) = %.10f\n", F)

	// Modifikasi

	fmt.Print("\nAFTER MODIFIKASI\n\n")

	fmt.Print("Nilai k = ")
	fmt.Scanln(&k)
	var Fakar2 float64 = 1
	for i := 0; i <= k; i++ {
		Fakar2 *= float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", Fakar2)
}

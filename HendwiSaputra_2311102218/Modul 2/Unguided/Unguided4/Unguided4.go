package main

import (
	"fmt"
)

func main() {
	var K int
	var F float64

	fmt.Print("Nilai K = ")
	fmt.Scanln(&K)

	F = float64((4*K+2)*(4*K+2)) / float64((4*K+1)*(4*K+3))
	fmt.Printf("Nilai f(K) = %.10f\n", F)

	fmt.Print("\nSetelah Modifikasi\n\n")

	fmt.Print("Nilai K = ")
	fmt.Scanln(&K)
	var faktorAkar float64 = 1
	for i := 0; i <= K; i++ {
		faktorAkar *= float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", faktorAkar)
}

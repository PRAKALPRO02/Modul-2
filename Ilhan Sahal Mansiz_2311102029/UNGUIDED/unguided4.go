package main

import (
	"fmt"
	"math"
)

func main(){
	// var k int
	
	// fmt.Print("Nilai K = ")
	// fmt.Scan(&k)

	// result := (math.Pow(float64(4*k+2),2) / float64(4*k+1)*float64(4*k+3))
	// fmt.Printf("Nilai f(k) = %10f\n", result)

	// Modifikasi
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	result := 1.0
	for i := 0; i < k; i++ {
		rumus := (math.Pow(float64(4*i+2), 2)) / (float64(4*i+1) * float64(4*i+3))

		if result * rumus > 0 {
			result *= rumus
		}
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
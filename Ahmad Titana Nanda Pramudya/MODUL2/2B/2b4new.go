package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var hasil float64

	
	for {
		fmt.Print("Nilai K  = ")
		fmt.Scan(&k)
		if k < 0 {
			break
		}

		hasil = math.Sqrt(2)

		fmt.Printf("Akar 2 (√2) = %.10f \n", hasil)
	}
}

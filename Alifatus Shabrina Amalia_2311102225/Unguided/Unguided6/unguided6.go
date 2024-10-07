package main

import (
	"fmt"
	"math"
)

func main() {
	var K int

	fmt.Print("Nilai K= ")
	fmt.Scanln(&K)

	akar2 := (1 + float64(K) + 3) / (1 + 2)

	fmt.Printf("Nilai akar2 = %.10f\n", akar2)
	fmt.Printf("Nilai akar 2 = %.10f\n", math.Sqrt(2))
}

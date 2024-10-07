package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var fk float64 = 1.0
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	for i := 0; i <= k; i++ {
		fk *= math.Pow(4*float64(i)+2, 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}
	fmt.Printf("nilai f(k) = %.10f", fk)
}

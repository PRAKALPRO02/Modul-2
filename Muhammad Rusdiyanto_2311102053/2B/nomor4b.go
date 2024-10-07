package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64
	var result float64 = 1
	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)
	for i := 0; i <= int(k); i++ {
		result *= math.Pow(4*float64(i)+2, 2) / float64(((4*i + 1) * (4*i + 3)))
	}
	fmt.Printf("Nilai akar 2 = %v\n", result)
}

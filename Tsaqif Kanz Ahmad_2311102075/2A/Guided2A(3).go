package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&radius)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	area := 4 * math.Pi * math.Pow(radius, 2)

	fmt.Printf("Bola dengan jari-jari %.2f memiliki volume %.4f dan luas permukaan %.4f\n", radius, volume, area)
}

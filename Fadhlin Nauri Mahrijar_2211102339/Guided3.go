package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari, volume, luas_kulit float64
	fmt.Print("Jejari = ")
	fmt.Scan(&jejari)
	volume = math.Pi * 4.0 / 3.0 * math.Pow(jejari, 3)
	luas_kulit = math.Pi * 4 * math.Pow(jejari, 2)
	fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f", jejari, volume, luas_kulit)
}

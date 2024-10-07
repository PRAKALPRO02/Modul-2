package main

import (
	"fmt"
	"math"
)

func main() {

	var jari, volume, luasKulit float64
	fmt.Print("Jejari = ")
	fmt.Scan(&jari)
	volume = math.Pi * (4.0 / 3.0) * math.Pow(jari, 3)
	luasKulit = 4 * math.Pi * math.Pow(jari, 2)
	fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f \n", jari, volume, luasKulit)
}

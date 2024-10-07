package main

import (
	"fmt"
	"math"
)

func main() {
	const phi = 3.1415926535
	var jari int
	fmt.Print("Jejari: ")
	fmt.Scan(&jari)

	volumeBola := (4.0 / 3.0) * phi * math.Pow(float64(jari), 3)
	luasBola := 4 * phi * math.Pow(float64(jari), 2)

	fmt.Println("Bola dengan jejari ", jari, "memiliki volume", volumeBola, "dan luas kulit", luasBola)
}

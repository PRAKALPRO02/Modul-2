package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari int
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari)

	volumeBola := (4.0 / 3.0) * pi * math.Pow(float64(jejari), 3)
	luasBola := 4 * pi * math.Pow(float64(jejari), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", jejari, volumeBola, luasBola)
}

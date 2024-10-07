package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("Jejari = ")
	fmt.Scan(&r)
	volume := 4.0 / 3.0 * math.Pi * math.Pow(r, 3)
	luas := 4 * math.Pi * math.Pow(r, 2)
	fmt.Println("Bola dengan jejari ", r, "memiliki volume", volume, "dan luas kulit", luas)
}

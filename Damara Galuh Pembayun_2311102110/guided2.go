package main

//Damara Galuh Pembayun 2311102110
import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	volumeBola := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
	luasBola := 4 * math.Pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volumeBola, luasBola)
}

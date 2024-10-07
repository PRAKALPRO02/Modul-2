package main

import "fmt"

func main() {
	var r, phi, volume, luas float64

	fmt.Print("Jejari : ")
	fmt.Scan(&r)

	phi = 3.1415926535
	volume = 4.0 / 3.0 * phi * r * r * r
	luas = 4 * phi * r * r

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}

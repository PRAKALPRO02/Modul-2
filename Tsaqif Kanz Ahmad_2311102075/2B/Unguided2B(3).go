package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2, totalBerat float64
	var oleng bool

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat = berat1 + berat2
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		oleng = math.Abs(berat1-berat2) >= 9
		fmt.Printf("Sepeda motor pak Andi oleng: %v\n", oleng)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var kiri_220, kanan_220 float32

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri_220, &kanan_220)

		if kiri_220 < 0 || kanan_220 < 0 || (kiri_220+kanan_220) > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(float64(kiri_220 - kanan_220))
		oleng := selisih >= 9

		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
}

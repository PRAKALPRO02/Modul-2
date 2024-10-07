package main

import "fmt"

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Kondisi berhenti:
		if berat1 < 0 || berat2 < 0 || berat1+berat2 > 150 {
			fmt.Println("Selesai.")
			break
		}

		// Hitung selisih berat dan tentukan apakah sepeda akan oleng
		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih // Ambil nilai absolut selisih
		}
		akanOleng := selisih >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)
	}
}

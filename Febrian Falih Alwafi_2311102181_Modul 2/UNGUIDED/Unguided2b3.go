package main

import "fmt"

func main() {
	var beratKiri, beratKanan float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		// Cek apakah total berat melebihi batas atau ada berat negatif
		if beratKiri < 0 || beratKanan < 0 || (beratKiri+beratKanan) > 150 {
			break
		}

		// Hitung apakah sepeda motor akan oleng
		sepedaMotorOleng := (beratKiri - beratKanan) >= 9 || (beratKanan - beratKiri) >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %v\n", sepedaMotorOleng)
	}

	fmt.Println("Proses selesai.")
}

package main

import (
	"fmt"
)

func main() {
	var warna1, warna2, warna3, warna4 string
	var berhasil bool = true

	fmt.Println("Masukkan warna percobaan (merah/kuning/hijau/ungu):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}

// Dwi Hesti Ariani_2311102094

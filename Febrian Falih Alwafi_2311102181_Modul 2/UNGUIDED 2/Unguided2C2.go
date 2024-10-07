package main

import "fmt"

func main() {
	for {
		var nilaiAkhir float64

		// Meminta input nilai akhir dari pengguna
		fmt.Print("Nilai akhir mata kuliah (masukkan -1 untuk keluar): ")
		fmt.Scanln(&nilaiAkhir)

		// Memeriksa apakah pengguna ingin keluar
		if nilaiAkhir == -1 {
			fmt.Println("Proses selesai.")
			break
		}

		// Menentukan grade berdasarkan nilai akhir
		nilaiHuruf := hitungNilaiHuruf(nilaiAkhir)

		// Menampilkan hasil grade
		fmt.Printf("Nilai mata kuliah: %v\n", nilaiHuruf)
	}
}

// Fungsi untuk menghitung nilai huruf berdasarkan nilai akhir
func hitungNilaiHuruf(nilai float64) string {
	if nilai > 80 {
		return "A"
	} else if nilai > 72.5 {
		return "AB"
	} else if nilai > 65 {
		return "B"
	} else if nilai > 57.5 {
		return "BC"
	} else if nilai > 50 {
		return "C"
	} else if nilai > 40 {
		return "D"
	} else {
		return "E"
	}
}

package main

import "fmt"

func main() {
	for {
		var nam float64
		var nmk string

		// Meminta input nilai akhir dari pengguna
		fmt.Print("Nilai akhir mata kuliah (masukkan -1 untuk keluar): ")
		fmt.Scanln(&nam)

		// Memeriksa apakah pengguna ingin keluar
		if nam == -1 {
			fmt.Println("Proses selesai.")
			break
		}

		// Menentukan grade berdasarkan nilai akhir
		if nam > 80 {
			nmk = "A"
		} else if nam > 72.5 {
			nmk = "AB"
		} else if nam > 65 {
			nmk = "B"
		} else if nam > 57.5 {
			nmk = "BC"
		} else if nam > 50 {
			nmk = "C"
		} else if nam > 40 {
			nmk = "D"
		} else {
			nmk = "E"
		}

		// Menampilkan hasil grade
		fmt.Printf("Nilai mata kuliah: %v\n", nmk)
	}
}

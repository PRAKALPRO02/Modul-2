package main

import "fmt"

func main() {
	var nilaiAkhir float64

	fmt.Print("Masukkan nilai akhir mata kuliah (0-100): ")
	fmt.Scan(&nilaiAkhir)

	var nilaiHuruf string

	switch {
	case nilaiAkhir >= 80:
		nilaiHuruf = "A"
	case nilaiAkhir >= 72.5:
		nilaiHuruf = "AB"
	case nilaiAkhir >= 65:
		nilaiHuruf = "B"
	case nilaiAkhir >= 57.5:
		nilaiHuruf = "BC"
	case nilaiAkhir >= 50:
		nilaiHuruf = "C"
	case nilaiAkhir >= 40:
		nilaiHuruf = "D"
	default:
		nilaiHuruf = "E"
	}

	fmt.Printf("Nilai huruf Anda adalah: %s\n", nilaiHuruf)
}

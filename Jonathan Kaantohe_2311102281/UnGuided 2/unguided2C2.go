package main

import "fmt"

func main() {
	var nilaiAkhir float64
	var nilaiHuruf string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nilaiAkhir)

	switch {
	case nilaiAkhir > 80:
		nilaiHuruf = "A"
	case nilaiAkhir > 72.5:
		nilaiHuruf = "AB"
	case nilaiAkhir > 65:
		nilaiHuruf = "B"
	case nilaiAkhir > 57.5:
		nilaiHuruf = "BC"
	case nilaiAkhir > 50:
		nilaiHuruf = "C"
	case nilaiAkhir > 40:
		nilaiHuruf = "D"
	default:
		nilaiHuruf = "E"
	}

	fmt.Println("Nilai mata kuliah:", nilaiHuruf)
}
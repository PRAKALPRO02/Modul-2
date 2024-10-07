package main

import "fmt"

func main() {
	var nilaiAkhir float64
	var grade string

	fmt.Print("Masukkan nilai akhir mata kuliah: ")
	fmt.Scanln(&nilaiAkhir)

	switch {
	case nilaiAkhir > 80 && nilaiAkhir <= 100:
		grade = "A"
	case nilaiAkhir > 72.5 && nilaiAkhir <= 80:
		grade = "AB"
	case nilaiAkhir > 65 && nilaiAkhir <= 72.5:
		grade = "B"
	case nilaiAkhir > 57.5 && nilaiAkhir <= 65:
		grade = "BC"
	case nilaiAkhir > 50 && nilaiAkhir <= 57.5:
		grade = "C"
	case nilaiAkhir > 40 && nilaiAkhir <= 50:
		grade = "D"
	case nilaiAkhir >= 0 && nilaiAkhir <= 40:
		grade = "E"
	default:
		grade = "Nilai tidak valid"
	}

	fmt.Println("Grade mata kuliah: ", grade)
}

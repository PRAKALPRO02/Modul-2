package main

import "fmt"

func main() {
	var nilaiAkhir float64
	var grade string

	fmt.Print("Masukkan nilai akhir mata kuliah: ")
	fmt.Scanln(&nilaiAkhir)

	if nilaiAkhir < 0 || nilaiAkhir > 100 {
		fmt.Println("Nilai harus antara 0 dan 100.")
		return
	}

	if nilaiAkhir >= 80 {
		grade = "A"
	} else if nilaiAkhir >= 72.5 {
		grade = "AB"
	} else if nilaiAkhir >= 65 {
		grade = "B"
	} else if nilaiAkhir >= 57.5 {
		grade = "BC"
	} else if nilaiAkhir >= 50 {
		grade = "C"
	} else if nilaiAkhir >= 40 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println("Nilai mata kuliah: ", grade)
}

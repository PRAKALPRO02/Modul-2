package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Meminta input nilai akhir
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	// Memastikan nilai berada dalam rentang yang valid (0-100)
	if nam < 0 || nam > 100 {
		fmt.Println("Nilai tidak valid! Harus antara 0 dan 100.")
		return
	}

	// Menghitung grade berdasarkan rentang nilai yang benar
	if nam >= 80 && nam <= 100 {
		nmk = "A"
	} else if nam >= 72.5 && nam < 80 {
		nmk = "AB"
	} else if nam >= 65 && nam < 72.5 {
		nmk = "B"
	} else if nam >= 57.5 && nam < 65 {
		nmk = "BC"
	} else if nam >= 50 && nam < 57.5 {
		nmk = "C"
	} else if nam >= 40 && nam < 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Menampilkan hasil grade
	fmt.Println("Nilai mata kuliah:", nmk)
}

package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float64 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float64
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam >= 80 {
		nmk = "A" // Jika nilai 80 atau lebih, indeksnya A
	} else if nam >= 72.5 {
		nmk = "AB" // Jika nilai 72.5 atau lebih, indeksnya B
	} else if nam >= 65 {
		nmk = "B" // Jika nilai 65 atau lebih, indeksnya C
	} else if nam >= 57.5 {
		nmk = "C" // Jika nilai 57.5 atau lebih, indeksnya D
	} else if nam >= 40 {
		nmk = "D" // Jika nilai 40 atau lebih, indeksnya E
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Println("Nilai mata kuliah: ", nmk)
}
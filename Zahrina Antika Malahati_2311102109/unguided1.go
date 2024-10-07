package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk memeriksa apakah urutan warna benar
func cekPercobaan(input [5][4]string) bool {
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if strings.ToLower(input[i][j]) != urutanBenar[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// Input percobaan
	var input [5][4]string

	// Masukkan data percobaan dari pengguna
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d : ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&input[i][j])
		}
	}

	// Memeriksa hasil percobaan
	if cekPercobaan(input) {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}

//Zahrina Antika Malahati_2311102109
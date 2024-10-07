package main

import (
	"fmt"
)

func main() {
	var eksperimen [5][4]string
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&eksperimen[i][j])
		}
	}

	// Pengecekan urutan
	berhasil := true
	for _, percobaan := range eksperimen {
		for i := range percobaan {
			if percobaan[i] != urutanBenar[i] {
				berhasil = false
				break
			}
		}
		if !berhasil {
			break
		}
	}

	fmt.Println("Urutan benar:", berhasil)
}

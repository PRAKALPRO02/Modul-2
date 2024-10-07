package main
//adam faizal 2311102088
import (
	"fmt"
)

func main() {
	// Urutan warna yang benar
	expected := [4]string{"merah", "kuning", "hijau", "ungu"}

	// Data percobaan dari input pengguna
	var percobaan [5][4]string
	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan percobaan %d : ", i+1)
		fmt.Scan(&percobaan[i][0], &percobaan[i][1], &percobaan[i][2], &percobaan[i][3])
	}

	// Periksa apakah semua percobaan sesuai dengan urutan warna yang diharapkan
	berhasil := true
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if percobaan[i][j] != expected[j] {
				berhasil = false
				break
			}
		}
		if !berhasil {
			break
		}
	}

	// Tampilkan hasil
	if berhasil {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}

package main

import "fmt"

func main() {
	// Array dengan urutan warna yang benar
	warnaTrue := [4]string{"merah", "kuning", "hijau", "ungu"}
	benar := true

	// Loop untuk 5 percobaan
	for i := 1; i <= 5; i++ {
		var inputWarna [4]string
		fmt.Printf("Percobaan %d: ", i)

		// Mengambil input untuk 4 warna
		fmt.Scan(&inputWarna[0], &inputWarna[1], &inputWarna[2], &inputWarna[3])

		// Mengecek setiap input warna apakah sesuai dengan urutan yang benar
		for j := 0; j < len(warnaTrue); j++ {
			if inputWarna[j] != warnaTrue[j] {
				benar = false
				break // Jika salah, langsung tandai salah dan lanjut ke percobaan berikutnya
			}
		}
	}

	// Menampilkan hasil akhir setelah 5 percobaan
	if benar {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}

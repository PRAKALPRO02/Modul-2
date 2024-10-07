package main

import "fmt"

func main() {
	var a, b, c, d string
	berhasil := true
	var i int // Inisialisasi variabel i

	for i < 5 { // Loop selama i kurang dari 6
		fmt.Printf("percobaan %d: ", i+1)
		fmt.Scanln(&a, &b, &c, &d)

		// Memeriksa apakah semua variabel tidak sesuai dengan kondisi yang diinginkan
		if a != "merah" && b != "kuning" && c != "hijau" && d != "ungu" {
			berhasil = false
		}

		i++ // Increment counter
	}

	// Menampilkan hasil
	fmt.Println(berhasil)
}

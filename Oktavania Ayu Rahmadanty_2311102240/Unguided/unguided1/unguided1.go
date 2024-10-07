package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var berhasil bool

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d:\n", i)
		fmt.Print("Warna gelas 1: ")
		fmt.Scan(&warna1)
		fmt.Print("Warna gelas 2: ")
		fmt.Scan(&warna2)
		fmt.Print("Warna gelas 3: ")
		fmt.Scan(&warna3)
		fmt.Print("Warna gelas 4: ")
		fmt.Scan(&warna4)

		if warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu" {
			berhasil = true
		} else {
			berhasil = false
			break // Jika salah satu percobaan gagal, langsung hentikan perulangan
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}
